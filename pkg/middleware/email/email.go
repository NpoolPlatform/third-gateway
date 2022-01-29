package email

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	code "github.com/NpoolPlatform/third-gateway/pkg/middleware/code"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func SendCode(ctx context.Context, in *npool.SendEmailCodeRequest) (*npool.SendEmailCodeResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	vCode := code.Generate6NumberCode()
	err = code.CreateCodeCache(ctx, &code.UserCode{
		AppID:       appID,
		Account:     in.GetEmailAddress(),
		AccountType: "email",
		UsedFor:     in.GetUsedFor(),
		Code:        vCode,
		ExpireAt:    time.Now().Add(10 * time.Minute),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create code cache: %v", err)
	}

	// TODO: get email template
	// TODO: send email

	return nil, nil
}

func VerifyCode(ctx context.Context, in *npool.VerifyEmailCodeRequest) (*npool.VerifyEmailCodeResponse, error) {
	return nil, nil
}
