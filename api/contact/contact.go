package contact

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/contact"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"
	contact "github.com/NpoolPlatform/third-middleware/pkg/client/contact"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ContactViaEmail(ctx context.Context, in *npool.ContactViaEmailRequest) (*npool.ContactViaEmailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateContact")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "CreateContact")

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
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.ContactViaEmailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ContactViaEmailResponse{}, nil
}
