package sms

import (
	"context"

	"github.com/google/uuid"

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

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	info, err := mgrcli.GetSMSTemplate(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	if info.GetAppID() != in.GetAppID() {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.PermissionDenied, "permission denied")
	}

	if _, err := uuid.Parse(in.GetTargetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.InvalidArgument, "LangID is invalid")
	}

	span = commontracer.TraceInvoker(span, "contact", "manager", "UpdateSMSTemplate")

	info, err = mgrcli.UpdateSMSTemplate(ctx, &mgrpb.SMSTemplateReq{
		ID:      &in.ID,
		LangID:  &in.TargetLangID,
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

func (s *Server) UpdateAppSMSTemplate(
	ctx context.Context,
	in *npool.UpdateAppSMSTemplateRequest,
) (
	*npool.UpdateAppSMSTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAppSMSTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateAppSMSTemplateResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	if _, err := uuid.Parse(in.GetTargetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.UpdateAppSMSTemplateResponse{}, status.Error(codes.InvalidArgument, "LangID is invalid")
	}

	span = commontracer.TraceInvoker(span, "contact", "manager", "UpdateSMSTemplate")

	info, err := mgrcli.UpdateSMSTemplate(ctx, &mgrpb.SMSTemplateReq{
		ID:      &in.ID,
		LangID:  &in.TargetLangID,
		Subject: &in.Subject,
		Message: &in.Message,
	})

	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateAppSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppSMSTemplateResponse{
		Info: info,
	}, nil
}
