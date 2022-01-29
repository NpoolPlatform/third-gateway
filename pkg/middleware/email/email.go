package email

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
)

func SendCode(ctx context.Context, in *npool.SendEmailCodeRequest) (*npool.SendEmailCodeResponse, error) {
	return nil, nil
}

func VerifyCode(ctx context.Context, in *npool.VerifyEmailCodeRequest) (*npool.VerifyEmailCodeResponse, error) {
	return nil, nil
}
