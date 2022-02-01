package appsmstemplate

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"

	"github.com/NpoolPlatform/third-gateway/pkg/db"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appsmstemplate"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateTemplate(info *npool.AppSMSTemplate) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetLangID()); err != nil {
		return xerrors.Errorf("invalid lang id: %v", err)
	}
	return nil
}

func dbRowToTemplate(row *ent.AppSMSTemplate) *npool.AppSMSTemplate {
	return &npool.AppSMSTemplate{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		LangID:  row.LangID.String(),
		UsedFor: row.UsedFor,
		Subject: row.Subject,
		Message: row.Message,
	}
}

func Create(ctx context.Context, in *npool.CreateAppSMSTemplateRequest) (*npool.CreateAppSMSTemplateResponse, error) {
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
		AppSMSTemplate.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetLangID(uuid.MustParse(in.GetInfo().GetLangID())).
		SetUsedFor(in.GetInfo().GetUsedFor()).
		SetSubject(in.GetInfo().GetSubject()).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app sms template: %v", err)
	}

	return &npool.CreateAppSMSTemplateResponse{
		Info: dbRowToTemplate(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppSMSTemplateRequest) (*npool.GetAppSMSTemplateResponse, error) {
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
		AppSMSTemplate.
		Query().
		Where(
			appsmstemplate.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app sms template: %v", err)
	}

	var template *npool.AppSMSTemplate
	for _, info := range infos {
		template = dbRowToTemplate(info)
		break
	}

	return &npool.GetAppSMSTemplateResponse{
		Info: template,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppSMSTemplateRequest) (*npool.UpdateAppSMSTemplateResponse, error) {
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
		AppSMSTemplate.
		UpdateOneID(id).
		SetSubject(in.GetInfo().GetSubject()).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app sms template: %v", err)
	}

	return &npool.UpdateAppSMSTemplateResponse{
		Info: dbRowToTemplate(info),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppSMSTemplatesByAppRequest) (*npool.GetAppSMSTemplatesByAppResponse, error) {
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
		AppSMSTemplate.
		Query().
		Where(
			appsmstemplate.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app sms template: %v", err)
	}

	templates := []*npool.AppSMSTemplate{}
	for _, info := range infos {
		templates = append(templates, dbRowToTemplate(info))
	}

	return &npool.GetAppSMSTemplatesByAppResponse{
		Infos: templates,
	}, nil
}

func GetByAppLangUsedFor(ctx context.Context, in *npool.GetAppSMSTemplateByAppLangUsedForRequest) (*npool.GetAppSMSTemplateByAppLangUsedForResponse, error) {
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
		AppSMSTemplate.
		Query().
		Where(
			appsmstemplate.And(
				appsmstemplate.AppID(appID),
				appsmstemplate.LangID(langID),
				appsmstemplate.UsedFor(in.GetUsedFor()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app sms template: %v", err)
	}

	var template *npool.AppSMSTemplate
	for _, info := range infos {
		template = dbRowToTemplate(info)
		break
	}

	return &npool.GetAppSMSTemplateByAppLangUsedForResponse{
		Info: template,
	}, nil
}
