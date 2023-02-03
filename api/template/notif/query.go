package notif

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/template/notif"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mgrpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"
	"github.com/NpoolPlatform/third-manager/pkg/client/template/notif"
)

func (s *Server) GetNotifTemplate(ctx context.Context, in *npool.GetNotifTemplateRequest) (*npool.GetNotifTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifTemplate")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "GetNotifTemplate")
	commontracer.TraceID(span, in.GetID())

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("validate", "ID", in.GetID())
		return &npool.GetNotifTemplateResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	info, err := notif.GetNotifTemplate(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.GetNotifTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifTemplateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetNotifTemplates(ctx context.Context, in *npool.GetNotifTemplatesRequest) (*npool.GetNotifTemplatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifTemplates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "GetNotifTemplates")

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID())
		return &npool.GetNotifTemplatesResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	infos, total, err := notif.GetNotifTemplates(ctx, &mgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
	}, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.GetNotifTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifTemplatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppNotifTemplates(
	ctx context.Context,
	in *npool.GetAppNotifTemplatesRequest,
) (
	*npool.GetAppNotifTemplatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppNotifTemplates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "contact", "manager", "GetAppNotifTemplates")
	commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "TargetAppID", in.GetTargetAppID())
		return &npool.GetAppNotifTemplatesResponse{}, status.Error(codes.InvalidArgument, "TargetAppID is invalid")
	}

	infos, total, err := notif.GetNotifTemplates(ctx, &mgrpb.Conds{
		AppID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
	}, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.GetAppNotifTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppNotifTemplatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
