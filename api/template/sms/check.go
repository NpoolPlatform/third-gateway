package sms

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/template/sms"
)

func validate(in *sms.CreateSMSTemplateRequest) error {
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

	if in.GetSubject() == "" {
		logger.Sugar().Errorw("validate", "Subject", in.GetSubject())
		return status.Error(codes.InvalidArgument, "Subject is empty")
	}
	if in.GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", in.GetMessage())
		return status.Error(codes.InvalidArgument, "Message is empty")
	}

	return nil
}
