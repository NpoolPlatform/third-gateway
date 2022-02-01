package appemailtemplate

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"

	"github.com/NpoolPlatform/third-gateway/pkg/db"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appemailtemplate"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateTemplate(info *npool.AppEmailTemplate) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetLangID()); err != nil {
		return xerrors.Errorf("invalid lang id: %v", err)
	}
	return nil
}

func dbRowToTemplate(row *ent.AppEmailTemplate) *npool.AppEmailTemplate {
	return &npool.AppEmailTemplate{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		LangID:   row.LangID.String(),
		UsedFor:  row.UsedFor,
		Sender:   row.Sender,
		ReplyTos: row.ReplyTos,
		CCTos:    row.CcTos,
		Subject:  row.Subject,
		Body:     row.Body,
	}
}

func Create(ctx context.Context, in *npool.CreateAppEmailTemplateRequest) (*npool.CreateAppEmailTemplateResponse, error) {
	if err := validateTemplate(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppEmailTemplate.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetLangID(uuid.MustParse(in.GetInfo().GetLangID())).
		SetUsedFor(in.GetInfo().GetUsedFor()).
		SetSender(in.GetInfo().GetSender()).
		SetReplyTos(in.GetInfo().GetReplyTos()).
		SetCcTos(in.GetInfo().GetCCTos()).
		SetSubject(in.GetInfo().GetSubject()).
		SetBody(in.GetInfo().GetBody()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app email template: %v", err)
	}

	return &npool.CreateAppEmailTemplateResponse{
		Info: dbRowToTemplate(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppEmailTemplateRequest) (*npool.GetAppEmailTemplateResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid template id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppEmailTemplate.
		Query().
		Where(
			appemailtemplate.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app email template: %v", err)
	}

	var template *npool.AppEmailTemplate
	for _, info := range infos {
		template = dbRowToTemplate(info)
		break
	}

	return &npool.GetAppEmailTemplateResponse{
		Info: template,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppEmailTemplateRequest) (*npool.UpdateAppEmailTemplateResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid template id: %v", err)
	}

	if err := validateTemplate(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppEmailTemplate.
		UpdateOneID(id).
		SetSender(in.GetInfo().GetSender()).
		SetReplyTos(in.GetInfo().GetReplyTos()).
		SetCcTos(in.GetInfo().GetCCTos()).
		SetSubject(in.GetInfo().GetSubject()).
		SetBody(in.GetInfo().GetBody()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app email template: %v", err)
	}

	return &npool.UpdateAppEmailTemplateResponse{
		Info: dbRowToTemplate(info),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppEmailTemplateByAppRequest) (*npool.GetAppEmailTemplateByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppEmailTemplate.
		Query().
		Where(
			appemailtemplate.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app email template: %v", err)
	}

	templates := []*npool.AppEmailTemplate{}
	for _, info := range infos {
		templates = append(templates, dbRowToTemplate(info))
	}

	return &npool.GetAppEmailTemplateByAppResponse{
		Infos: templates,
	}, nil
}

func GetByAppLangUsedFor(ctx context.Context, in *npool.GetAppEmailTemplateByAppLangUsedForRequest) (*npool.GetAppEmailTemplateByAppLangUsedForResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	langID, err := uuid.Parse(in.GetLangID())
	if err != nil {
		return nil, xerrors.Errorf("invalid lang id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppEmailTemplate.
		Query().
		Where(
			appemailtemplate.And(
				appemailtemplate.AppID(appID),
				appemailtemplate.LangID(langID),
				appemailtemplate.UsedFor(in.GetUsedFor()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app email template: %v", err)
	}

	var template *npool.AppEmailTemplate
	for _, info := range infos {
		template = dbRowToTemplate(info)
		break
	}

	return &npool.GetAppEmailTemplateByAppLangUsedForResponse{
		Info: template,
	}, nil
}
