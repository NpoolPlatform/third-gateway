package email

import (
	"context"
	"strings"

	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"
	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	contactcrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appcontact"
	templatecrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appemailtemplate"
	grpc2 "github.com/NpoolPlatform/third-gateway/pkg/grpc"
	code "github.com/NpoolPlatform/third-gateway/pkg/middleware/code"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func SendCode(ctx context.Context, in *npool.SendEmailCodeRequest) (*npool.SendEmailCodeResponse, error) {
	template, err := templatecrud.GetByAppLangUsedFor(ctx, &npool.GetAppEmailTemplateByAppLangUsedForRequest{
		AppID:   in.GetAppID(),
		LangID:  in.GetLangID(),
		UsedFor: in.GetUsedFor(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app email template: %v", err)
	}
	if template.Info == nil {
		return nil, xerrors.Errorf("fail get app email template")
	}

	body, err := buildBody(ctx, in, template.Info.Body)
	if err != nil {
		return nil, xerrors.Errorf("fail build email body: %v", err)
	}

	if in.GetToUsername() == "" {
		body = strings.ReplaceAll(body, NameTemplate, template.Info.DefaultToUsername)
	}

	err = sendEmailByAWS(template.Info.Subject, body, template.Info.Sender, in.GetEmailAddress())
	if err != nil {
		return nil, xerrors.Errorf("fail send email: %v", err)
	}

	return &npool.SendEmailCodeResponse{}, nil
}

func VerifyCode(ctx context.Context, in *npool.VerifyEmailCodeRequest) (*npool.VerifyEmailCodeResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	emailAddr := in.GetEmailAddress()

	_, err = uuid.Parse(in.GetUserID())
	if err == nil {
		resp, err := grpc2.GetAppUserByAppUser(ctx, &appusermgrpb.GetAppUserByAppUserRequest{
			AppID:  in.GetAppID(),
			UserID: in.GetUserID(),
		})
		if err != nil {
			return nil, xerrors.Errorf("invalid app user: %v", err)
		}
		emailAddr = resp.Info.EmailAddress
	}

	userCode := code.UserCode{
		AppID:       appID,
		Account:     emailAddr,
		AccountType: AccountType,
		UsedFor:     in.GetUsedFor(),
		Code:        in.GetCode(),
	}

	err = code.VerifyCodeCache(ctx, &userCode)
	if err != nil {
		return nil, xerrors.Errorf("invalid code: %v", err)
	}

	return &npool.VerifyEmailCodeResponse{}, nil
}

func Contact(ctx context.Context, in *npool.ContactRequest) (*npool.ContactResponse, error) {
	resp, err := contactcrud.GetByAppUsedForAccountType(ctx, &npool.GetAppContactByAppUsedForAccountTypeRequest{
		AppID:       in.GetAppID(),
		UsedFor:     in.GetUsedFor(),
		AccountType: AccountType,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app contact: %v", err)
	}

	err = sendEmailByAWS(in.GetSubject(), in.GetBody(), in.GetSender(), resp.Info.Account)
	if err != nil {
		return nil, xerrors.Errorf("fail send email: %v", err)
	}

	return &npool.ContactResponse{}, nil
}
