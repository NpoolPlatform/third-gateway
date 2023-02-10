package notif

import (
	"context"
	"fmt"

	usedfor "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/template/notif"

	mgrpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/template/notif"
	mgrcli "github.com/NpoolPlatform/third-manager/pkg/client/template/notif"

	appusermgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/app"

	applangmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

//nolint
func validate(ctx context.Context, in *notif.CreateNotifTemplateRequest) error {
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
		return fmt.Errorf("applang not exist")
	}

	switch in.GetUsedFor() {
	case usedfor.EventType_WithdrawalRequest:
	case usedfor.EventType_WithdrawalCompleted:
	case usedfor.EventType_DepositReceived:
	case usedfor.EventType_KYCApproved:
	case usedfor.EventType_KYCRejected:
	case usedfor.EventType_Announcement:
	default:
		return fmt.Errorf("EventType is invalid")
	}

	if in.GetTitle() == "" {
		logger.Sugar().Errorw("validate", "Title", in.GetTitle())
		return status.Error(codes.InvalidArgument, "Title is empty")
	}
	if in.GetContent() == "" {
		logger.Sugar().Errorw("validate", "Content", in.GetContent())
		return status.Error(codes.InvalidArgument, "Content is empty")
	}
	if in.GetSender() == "" {
		logger.Sugar().Errorw("validate", "Sender", in.GetSender())
		return status.Error(codes.InvalidArgument, "Sender is empty")
	}
	exist, err = mgrcli.ExistNotifTemplateConds(ctx, &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
		UsedFor: &commonpb.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(in.GetUsedFor()),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return status.Error(codes.Internal, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("validate", "Notif template already exists")
		return status.Error(codes.AlreadyExists, "Notif template already exists")
	}

	return nil
}
