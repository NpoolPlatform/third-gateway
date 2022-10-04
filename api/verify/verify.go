package verify

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/verify"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mverify "github.com/NpoolPlatform/third-gateway/pkg/verify"
)

//nolint:gocyclo
func (s *Server) SendCode(ctx context.Context, in *npool.SendCodeRequest) (*npool.SendCodeResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateContact")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("SendCode", "AppID", in.GetAppID())
		return &npool.SendCodeResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	switch in.GetUsedFor() {
	case usedfor.UsedFor_Contact:
		fallthrough //nolint
	case usedfor.UsedFor_SetWithdrawAddress:
		fallthrough //nolint
	case usedfor.UsedFor_Withdraw:
		fallthrough //nolint
	case usedfor.UsedFor_CreateInvitationCode:
		fallthrough //nolint
	case usedfor.UsedFor_SetCommission:
		fallthrough //nolint
	case usedfor.UsedFor_SetTransferTargetUser:
		fallthrough //nolint
	case usedfor.UsedFor_Transfer:
		if _, err := uuid.Parse(in.GetUserID()); err != nil {
			logger.Sugar().Errorw("SendCode", "UserID", in.GetUserID())
			return &npool.SendCodeResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
		}
	}

	if in.Account == nil && in.UserID == nil {
		logger.Sugar().Errorw("SendCode", "Account", in.GetAccount(), "UserID", in.GetUserID())
		return &npool.SendCodeResponse{}, status.Error(codes.InvalidArgument, "Account and UserID cannot all be empty")
	}

	switch in.GetAccountType() {
	case signmethod.SignMethodType_Email:
	case signmethod.SignMethodType_Mobile:
	default:
		logger.Sugar().Errorw("SendCode", "AccountType", in.GetAccountType())
		return &npool.SendCodeResponse{}, status.Error(codes.InvalidArgument, "AccountType is invalid")
	}

	usedFor := false
	for key := range usedfor.UsedFor_value {
		if in.GetUsedFor().String() == key {
			usedFor = true
		}
	}
	if !usedFor {
		logger.Sugar().Errorw("SendCode", "UsedFor", in.GetUsedFor())
		return &npool.SendCodeResponse{}, status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}

	span = commontracer.TraceInvoker(span, "contact", "manager", "CreateEmailTemplate")

	err = mverify.SendCode(
		ctx,
		in.GetAppID(),
		in.GetLangID(),
		in.UserID,
		in.Account,
		in.GetAccountType(),
		in.GetUsedFor(),
		in.ToUsername,
	)
	if err != nil {
		logger.Sugar().Errorw("SendCode", "err", err)
		return &npool.SendCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.SendCodeResponse{}, nil
}
