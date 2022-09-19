package contact

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/contact"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
)

func validate(info *contact.CreateContactRequest) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID())
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	switch info.GetUsedFor() {
	case usedfor.UsedFor_Contact:
	default:
		logger.Sugar().Errorw("validate", "UsedFor", info.GetUsedFor())
		return status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}

	if info.GetAccount() == "" {
		logger.Sugar().Errorw("validate", "Account", info.GetAccount())
		return status.Error(codes.InvalidArgument, "Account is empty")
	}

	switch info.GetAccountType() {
	case signmethod.SignMethodType_Email:
	case signmethod.SignMethodType_Mobile:
	default:
		logger.Sugar().Errorw("validate", "AccountType", info.GetAccountType())
		return status.Error(codes.InvalidArgument, "AccountType is invalid")
	}

	if info.GetSender() == "" {
		logger.Sugar().Errorw("validate", "Sender", info.GetSender())
		return status.Error(codes.InvalidArgument, "Sender is empty")
	}

	return nil
}
