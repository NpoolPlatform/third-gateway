package notif

import (
	"context"

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
)

func (s *Server) CreateNotifTemplate(
	ctx context.Context,
	in *npool.CreateNotifTemplateRequest,
) (
	*npool.CreateNotifTemplateResponse,
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

	err = validate(ctx, in)
	if err != nil {
		return nil, err
	}

	span = commontracer.TraceInvoker(span, "contact", "manager", "CreateNotifTemplate")

	info, err := mgrcli.CreateNotifTemplate(ctx, &mgrpb.NotifTemplateReq{
		AppID:   &in.AppID,
		LangID:  &in.TargetLangID,
		UsedFor: &in.UsedFor,
		Title:   &in.Title,
		Content: &in.Content,
		Sender:  &in.Sender,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.CreateNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAppNotifTemplate(
	ctx context.Context,
	in *npool.CreateAppNotifTemplateRequest,
) (
	*npool.CreateAppNotifTemplateResponse,
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

	span = commontracer.TraceInvoker(span, "contact", "manager", "CreateNotifTemplate")

	err = validate(ctx, &npool.CreateNotifTemplateRequest{
		AppID:        in.TargetAppID,
		TargetLangID: in.TargetLangID,
		UsedFor:      in.UsedFor,
		Title:        in.Title,
		Content:      in.Content,
		Sender:       in.Sender,
	})
	if err != nil {
		return nil, err
	}

	info, err := mgrcli.CreateNotifTemplate(ctx, &mgrpb.NotifTemplateReq{
		AppID:   &in.TargetAppID,
		LangID:  &in.TargetLangID,
		UsedFor: &in.UsedFor,
		Title:   &in.Title,
		Content: &in.Content,
		Sender:  &in.Sender,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.CreateAppNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppNotifTemplateResponse{
		Info: info,
	}, nil
}
