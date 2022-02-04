package sms

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	templatecrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appsmstemplate"
	code "github.com/NpoolPlatform/third-gateway/pkg/middleware/code"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func SendCode(ctx context.Context, in *npool.SendSMSCodeRequest) (*npool.SendSMSCodeResponse, error) {
	template, err := templatecrud.GetByAppLangUsedFor(ctx, &npool.GetAppSMSTemplateByAppLangUsedForRequest{
		AppID:   in.GetAppID(),
		LangID:  in.GetLangID(),
		UsedFor: in.GetUsedFor(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app sms template: %v", err)
	}
	if template.Info == nil {
		return nil, xerrors.Errorf("fail get app sms template")
	}

	body, err := buildBody(ctx, in, template.Info.Message)
	if err != nil {
		return nil, xerrors.Errorf("fail build sms body: %v", err)
	}

	err = sendSMSByAWS(body, in.GetPhoneNO())
	if err != nil {
		return nil, xerrors.Errorf("fail send sms: %v", err)
	}

	return &npool.SendSMSCodeResponse{}, nil
}

func VerifyCode(ctx context.Context, in *npool.VerifySMSCodeRequest) (*npool.VerifySMSCodeResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userCode := code.UserCode{
		AppID:       appID,
		Account:     in.GetPhoneNO(),
		AccountType: AccountType,
		UsedFor:     in.GetUsedFor(),
		Code:        in.GetCode(),
	}

	err = code.VerifyCodeCache(ctx, &userCode)
	if err != nil {
		return nil, xerrors.Errorf("invalid code: %v", err)
	}

	return &npool.VerifySMSCodeResponse{}, nil
}
