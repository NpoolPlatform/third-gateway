package verify

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"

	usermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"

	"github.com/NpoolPlatform/third-middleware/pkg/client/verify"
)

func SendCode(
	ctx context.Context,
	appID,
	langID string,
	userID,
	account *string,
	accountType signmethod.SignMethodType,
	usedFor usedfor.UsedFor,
	toUserName *string,
) error {
	switch usedFor {
	case usedfor.UsedFor_Contact:
		fallthrough //nolint
	case usedfor.UsedFor_SetWithdrawAddress:
		fallthrough //nolint
	case usedfor.UsedFor_Withdraw:
		fallthrough //nolint
	case usedfor.UsedFor_CreateInvitationCode:
		fallthrough //nolint
	case usedfor.UsedFor_SetCommission:
		fallthrough //nolint
	case usedfor.UsedFor_SetTransferTargetUser:
		fallthrough //nolint
	case usedfor.UsedFor_Transfer:
		if userID != nil && *userID != "" {
			user, err := usermwcli.GetUser(ctx, appID, *userID)
			if err != nil {
				return err
			}
			switch accountType {
			case signmethod.SignMethodType_Mobile:
				account = &user.PhoneNO
			case signmethod.SignMethodType_Email:
				account = &user.EmailAddress
			}
		}
	}

	if *account == "" {
		return fmt.Errorf("invalid account")
	}

	err := verify.SendCode(ctx, appID, langID, *account, accountType, usedFor, toUserName)
	if err != nil {
		return err
	}

	return nil
}
