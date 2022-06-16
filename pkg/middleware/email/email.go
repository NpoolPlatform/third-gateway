package email

import (
	"context"
	"fmt"
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
		if err != nil || resp.Info == nil {
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

func Contact(ctx context.Context, in *npool.ContactByEmailRequest) (*npool.ContactByEmailResponse, error) {
	resp, err := contactcrud.GetByAppUsedForAccountType(ctx, &npool.GetAppContactByAppUsedForAccountTypeRequest{
		AppID:       in.GetAppID(),
		UsedFor:     in.GetUsedFor(),
		AccountType: AccountType,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app contact: %v", err)
	}

	body := fmt.Sprintf("From: %v<br>Name: %v<br>%v", in.GetSender(), in.GetSenderName(), in.GetBody())
	body = strings.ReplaceAll(body, "\n", "<br>")

	err = sendEmailByAWS(in.GetSubject(), body, resp.Info.Sender, resp.Info.Account, in.GetSender())
	if err != nil {
		return nil, xerrors.Errorf("fail send email: %v", err)
	}

	return &npool.ContactByEmailResponse{}, nil
}

func Notify(ctx context.Context, in *npool.NotifyEmailRequest) (*npool.NotifyEmailResponse, error) {
	to, err := grpc2.GetAppUserByAppUser(ctx, &appusermgrpb.GetAppUserByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetReceiverID(),
	})
	if err != nil || to.Info == nil {
		return nil, xerrors.Errorf("fail get user: %v", err)
	}
	if to.Info.EmailAddress == "" {
		return &npool.NotifyEmailResponse{}, nil
	}

	from, err := grpc2.GetAppUserByAppUser(ctx, &appusermgrpb.GetAppUserByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil || from.Info == nil {
		return nil, xerrors.Errorf("fail get user: %v", err)
	}

	template, err := templatecrud.GetByAppLangUsedFor(ctx, &npool.GetAppEmailTemplateByAppLangUsedForRequest{
		AppID:   in.GetAppID(),
		LangID:  in.GetLangID(),
		UsedFor: in.GetUsedFor(),
	})
	if err != nil || template.Info == nil {
		return &npool.NotifyEmailResponse{}, nil
	}

	body := strings.ReplaceAll(template.Info.Body, "{ FROM }", from.Info.EmailAddress)
	body = strings.ReplaceAll(body, "{ TO }", to.Info.EmailAddress)
	body = strings.ReplaceAll(body, "{ RECEIVER }", in.GetReceiverName())
	body = strings.ReplaceAll(body, "{ SENDER }", in.GetSenderName())

	err = sendEmailByAWS(template.Info.Subject, body, template.Info.Sender, to.Info.EmailAddress)
	if err != nil {
		return nil, xerrors.Errorf("fail send email: %v", err)
	}

	return &npool.NotifyEmailResponse{}, nil
}
