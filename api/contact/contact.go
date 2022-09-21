package contact

import (
	"context"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/contact"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	contact "github.com/NpoolPlatform/third-middleware/pkg/client/contact"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ContactViaEmail(ctx context.Context, in *npool.ContactViaEmailRequest) (*npool.ContactViaEmailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ContactViaEmail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("ContactViaEmail", "AppID", in.GetAppID())
		return &npool.ContactViaEmailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	switch in.GetUsedFor() {
	case usedfor.UsedFor_Contact:
	default:
		logger.Sugar().Errorw("ContactViaEmail", "UsedFor", in.GetUsedFor())
		return &npool.ContactViaEmailResponse{}, status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}

	if in.GetSender() == "" {
		logger.Sugar().Errorw("ContactViaEmail", "Sender", in.GetSender())
		return &npool.ContactViaEmailResponse{}, status.Error(codes.InvalidArgument, "Sender is empty")
	}
	if in.GetSubject() == "" {
		logger.Sugar().Errorw("ContactViaEmail", "Subject", in.GetSubject())
		return &npool.ContactViaEmailResponse{}, status.Error(codes.InvalidArgument, "Subject is empty")
	}
	if in.GetBody() == "" {
		logger.Sugar().Errorw("ContactViaEmail", "Body", in.GetBody())
		return &npool.ContactViaEmailResponse{}, status.Error(codes.InvalidArgument, "Body is empty")
	}

	span = commontracer.TraceInvoker(span, "contact", "middleware", "ContactViaEmail")

	err = contact.ContactViaEmail(
		ctx,
		in.GetAppID(),
		in.GetUsedFor(),
		in.GetSender(),
		in.GetSubject(),
		in.GetBody(),
		in.GetSenderName(),
	)
	if err != nil {
		logger.Sugar().Errorw("ContactViaEmail", "err", err)
		return &npool.ContactViaEmailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ContactViaEmailResponse{}, nil
}
