package sms

import (
	"context"
	"strings"
	"time"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	constant "github.com/NpoolPlatform/third-gateway/pkg/const"
	code "github.com/NpoolPlatform/third-gateway/pkg/middleware/code"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	CodeTemplate = "{{ CODE }}"
)

func buildWithCode(ctx context.Context, in *npool.SendSMSCodeRequest, template string) (string, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return "", xerrors.Errorf("invalid app id: %v", err)
	}

	vCode := code.Generate6NumberCode()
	userCode := code.UserCode{
		AppID:       appID,
		Account:     in.GetPhoneNO(),
		AccountType: AccountType,
		UsedFor:     in.GetUsedFor(),
		Code:        vCode,
		NextAt:      time.Now().Add(1 * time.Minute),
		ExpireAt:    time.Now().Add(10 * time.Minute),
	}
	if ok, err := code.Nextable(ctx, &userCode); err != nil || !ok {
		return "", xerrors.Errorf("wait for next code generation")
	}

	err = code.CreateCodeCache(ctx, &userCode)
	if err != nil {
		return "", xerrors.Errorf("fail create code cache: %v", err)
	}

	return strings.ReplaceAll(template, CodeTemplate, vCode), nil
}

func buildBody(ctx context.Context, in *npool.SendSMSCodeRequest, template string) (string, error) {
	switch in.GetUsedFor() {
	case constant.UsedForSignup:
		fallthrough // nolint
	case constant.UsedForSignin:
		fallthrough // nolint
	case constant.UsedForSetWithdrawAddress:
		fallthrough // nolint
	case constant.UsedForWithdraw:
		fallthrough //nolint
	case constant.UsedForUpdate:
		return buildWithCode(ctx, in, template)
	case constant.UsedForContact:
		return "", xerrors.Errorf("NOT IMPLEMENTED")
	}

	return "", xerrors.Errorf("invalid used for")
}
