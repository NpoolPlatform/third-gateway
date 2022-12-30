package sms

import (
	"context"
	"fmt"

	appusermgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/template/sms"

	mgrpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/sms"
	mgrcli "github.com/NpoolPlatform/third-manager/pkg/client/template/sms"

	applangmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func validate(ctx context.Context, in *sms.CreateSMSTemplateRequest) error {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID())
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	exist, err := appusermgrcli.ExistApp(ctx, in.GetAppID())
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID())
		return status.Error(codes.InvalidArgument, "AppID is not exist")
	}

	if _, err := uuid.Parse(in.GetTargetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "TargetLangID", in.GetTargetLangID())
		return status.Error(codes.InvalidArgument, "TargetLangID is invalid")
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
		return err
	}
	if appLang == nil {
		logger.Sugar().Errorw("validate", "err", err)
		return fmt.Errorf("applang not exist")
	}

	usedFor := false
	for key := range usedfor.UsedFor_value {
		if key == in.UsedFor.String() && in.UsedFor != usedfor.UsedFor_DefaultUsedFor {
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

	exist, err = mgrcli.ExistSMSTemplateConds(ctx, &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
		UsedFor: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(in.GetUsedFor().Number()),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return status.Error(codes.Internal, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("validate", "SMS template already exists")
		return status.Error(codes.AlreadyExists, "SMS template already exists")
	}

	return nil
}
