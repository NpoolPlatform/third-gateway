package email

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/template/email"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mgrpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/email"
	mgrcli "github.com/NpoolPlatform/third-manager/pkg/client/template/email"
)

func (s *Server) CreateEmailTemplate(
	ctx context.Context,
	in *npool.CreateEmailTemplateRequest,
) (
	*npool.CreateEmailTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateContact")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "CreateEmailTemplate")

	err = validate(ctx, in)
	if err != nil {
		return nil, err
	}

	info, err := mgrcli.CreateEmailTemplate(ctx, &mgrpb.EmailTemplateReq{
		AppID:             &in.AppID,
		LangID:            &in.LangID,
		UsedFor:           &in.UsedFor,
		Sender:            &in.Sender,
		ReplyTos:          in.ReplyTos,
		CCTos:             in.CCTos,
		Subject:           &in.Subject,
		Body:              &in.Body,
		DefaultToUsername: &in.DefaultToUsername,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.CreateEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEmailTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAppEmailTemplate(
	ctx context.Context,
	in *npool.CreateAppEmailTemplateRequest,
) (
	*npool.CreateAppEmailTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateContact")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "CreateEmailTemplate")

	err = validate(ctx, &npool.CreateEmailTemplateRequest{
		AppID:             in.TargetAppID,
		LangID:            in.LangID,
		UsedFor:           in.UsedFor,
		Sender:            in.Sender,
		ReplyTos:          in.ReplyTos,
		CCTos:             in.CCTos,
		Subject:           in.Subject,
		Body:              in.Body,
		DefaultToUsername: in.DefaultToUsername,
	})
	if err != nil {
		return nil, err
	}

	info, err := mgrcli.CreateEmailTemplate(ctx, &mgrpb.EmailTemplateReq{
		AppID:             &in.TargetAppID,
		LangID:            &in.LangID,
		UsedFor:           &in.UsedFor,
		Sender:            &in.Sender,
		ReplyTos:          in.ReplyTos,
		CCTos:             in.CCTos,
		Subject:           &in.Subject,
		Body:              &in.Body,
		DefaultToUsername: &in.DefaultToUsername,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.CreateAppEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppEmailTemplateResponse{
		Info: info,
	}, nil
}
