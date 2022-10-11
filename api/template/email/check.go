package email

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/template/email"

	mgrpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/email"
	mgrcli "github.com/NpoolPlatform/third-manager/pkg/client/template/email"

	appusermgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/app"

	internationalizationcli "github.com/NpoolPlatform/internationalization/pkg/client/lang"
)

//nolint
func validate(ctx context.Context, in *email.CreateEmailTemplateRequest) error {
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

	exist, err = internationalizationcli.ExistLang(ctx, in.GetTargetLangID())
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return status.Error(codes.Internal, err.Error())
	}

	if !exist {
		logger.Sugar().Errorw("validate", "TargetLangID", in.GetTargetLangID())
		return status.Error(codes.InvalidArgument, "TargetLangID is not exist")
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

	if in.GetSender() == "" {
		logger.Sugar().Errorw("validate", "Sender", in.GetSender())
		return status.Error(codes.InvalidArgument, "Sender is empty")
	}
	if in.GetSubject() == "" {
		logger.Sugar().Errorw("validate", "Subject", in.GetSubject())
		return status.Error(codes.InvalidArgument, "Subject is empty")
	}
	if in.GetDefaultToUsername() == "" {
		logger.Sugar().Errorw("validate", "DefaultToUsername", in.GetDefaultToUsername())
		return status.Error(codes.InvalidArgument, "DefaultToUsername is empty")
	}

	exist, err = mgrcli.ExistEmailTemplateConds(ctx, &mgrpb.Conds{
		AppID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		LangID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
		UsedFor: &npool.Int32Val{
			Op:    cruder.EQ,
			Value: int32(in.GetUsedFor().Number()),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return status.Error(codes.Internal, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("validate", "Email template already exists")
		return status.Error(codes.AlreadyExists, "Email template already exists")
	}

	return nil
}
