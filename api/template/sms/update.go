package sms

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/template/sms"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mgrpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/sms"
	mgrcli "github.com/NpoolPlatform/third-manager/pkg/client/template/sms"
)

func (s *Server) UpdateSMSTemplate(
	ctx context.Context,
	in *npool.UpdateSMSTemplateRequest,
) (
	*npool.UpdateSMSTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSMSTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "UpdateSMSTemplate")

	info, err := mgrcli.UpdateSMSTemplate(ctx, &mgrpb.SMSTemplateReq{
		ID:      &in.ID,
		Subject: &in.Subject,
		Message: &in.Message,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSMSTemplateResponse{
		Info: info,
	}, nil
}
