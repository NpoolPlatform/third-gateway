package notif

import (
	"context"

	usedfor "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/template/notif"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mgrpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"
	mgrcli "github.com/NpoolPlatform/third-manager/pkg/client/template/notif"

	applangmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

//nolint:funlen,gocyclo
func (s *Server) UpdateNotifTemplate(
	ctx context.Context,
	in *npool.UpdateNotifTemplateRequest,
) (
	*npool.UpdateNotifTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateNotifTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	if _, err := uuid.Parse(in.GetTargetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "LangID is invalid")
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "LangID is invalid")
	}

	if in.GetTitle() == "" && in.Title != nil {
		logger.Sugar().Errorw("validate", "Title", in.GetTitle())
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "Title is empty")
	}

	if in.GetContent() == "" && in.Content != nil {
		logger.Sugar().Errorw("validate", "Content", in.GetContent())
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "Content is empty")
	}

	info, err := mgrcli.GetNotifTemplate(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	if info.GetAppID() != in.GetAppID() {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	if in.UsedFor != nil {
		switch in.GetUsedFor() {
		case usedfor.EventType_WithdrawalRequest:
		case usedfor.EventType_WithdrawalCompleted:
		case usedfor.EventType_DepositReceived:
		case usedfor.EventType_KYCApproved:
		case usedfor.EventType_KYCRejected:
		default:
			logger.Sugar().Errorw("validate", "err", err)
			return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.PermissionDenied, "UsedFor is invalid")
		}
	}

	appLang, err := applangmwcli.GetLangOnly(ctx, &applangmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	if appLang == nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, "AppLang not exist")
	}

	span = commontracer.TraceInvoker(span, "contact", "manager", "UpdateNotifTemplate")

	info, err = mgrcli.UpdateNotifTemplate(ctx, &mgrpb.NotifTemplateReq{
		ID:      &in.ID,
		AppID:   &in.AppID,
		LangID:  in.TargetLangID,
		UsedFor: in.UsedFor,
		Title:   in.Title,
		Content: in.Content,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNotifTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateAppNotifTemplate(
	ctx context.Context,
	in *npool.UpdateAppNotifTemplateRequest,
) (
	*npool.UpdateAppNotifTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateNotifTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "UpdateNotifTemplate")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateAppNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	if _, err := uuid.Parse(in.GetTargetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateAppNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "LangID is invalid")
	}

	appLang, err := applangmwcli.GetLangOnly(ctx, &applangmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateAppNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	if appLang == nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateAppNotifTemplateResponse{}, status.Error(codes.Internal, "AppLang not exist")
	}

	info, err := mgrcli.UpdateNotifTemplate(ctx, &mgrpb.NotifTemplateReq{
		ID:      &in.ID,
		AppID:   &in.TargetAppID,
		LangID:  in.TargetLangID,
		Title:   in.Title,
		Content: in.Content,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateAppNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppNotifTemplateResponse{
		Info: info,
	}, nil
}
