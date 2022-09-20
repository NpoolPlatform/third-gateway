package email

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/template/email"
)

func validate(in *email.CreateEmailTemplateRequest) error {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID())
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(in.GetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "LangID", in.GetLangID())
		return status.Error(codes.InvalidArgument, "LangID is invalid")
	}

	usedFor := false
	for key := range usedfor.UsedFor_value {
		if key == in.UsedFor.String() {
			usedFor = true
		}
	}

	if !usedFor {
		logger.Sugar().Errorw("validate", "UsedFor", in.GetUsedFor())
		return status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}

	if in.GetSender() == "" {
		logger.Sugar().Errorw("validate", "Sender", in.GetSender())
		return status.Error(codes.InvalidArgument, "Sender is empty")
	}
	if in.GetSubject() == "" {
		logger.Sugar().Errorw("validate", "Subject", in.GetSubject())
		return status.Error(codes.InvalidArgument, "Subject is empty")
	}
	if in.GetBody() == "" {
		logger.Sugar().Errorw("validate", "Body", in.GetBody())
		return status.Error(codes.InvalidArgument, "Body is empty")
	}
	if in.GetDefaultToUsername() == "" {
		logger.Sugar().Errorw("validate", "DefaultToUsername", in.GetDefaultToUsername())
		return status.Error(codes.InvalidArgument, "DefaultToUsername is empty")
	}

	return nil
}
