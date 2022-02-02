package email

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	templatecrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appemailtemplate"

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

	err = sendEmailByAWS(template.Info.Subject, body, template.Info.Sender, in.GetEmailAddress())
	if err != nil {
		return nil, xerrors.Errorf("fail send email: %v", err)
	}

	return &npool.SendEmailCodeResponse{}, nil
}

func VerifyCode(ctx context.Context, in *npool.VerifyEmailCodeRequest) (*npool.VerifyEmailCodeResponse, error) {
	return nil, nil
}
