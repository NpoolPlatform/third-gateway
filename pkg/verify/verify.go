package verify

import (
	"context"

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
) error {
	var toUserName *string

	if userID != nil {
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
		if user.GetUsername() != "" {
			toUserName = &user.Username
		}
	}

	err := verify.SendCode(ctx, appID, langID, *account, accountType, usedFor, toUserName)
	if err != nil {
		return err
	}

	return nil
}
