package sms

import (
	"context"

	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"
	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	templatecrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appsmstemplate"
	grpc2 "github.com/NpoolPlatform/third-gateway/pkg/grpc"
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

	phoneNO := in.GetPhoneNO()

	_, err = uuid.Parse(in.GetUserID())
	if err == nil {
		resp, err := grpc2.GetAppUserByAppUser(ctx, &appusermgrpb.GetAppUserByAppUserRequest{
			AppID:  in.GetAppID(),
			UserID: in.GetUserID(),
		})
		if err != nil {
			return nil, xerrors.Errorf("invalid app user: %v", err)
		}
		phoneNO = resp.Info.PhoneNO
	}

	userCode := code.UserCode{
		AppID:       appID,
		Account:     phoneNO,
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
