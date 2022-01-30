package appemailtemplate

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"

	"github.com/NpoolPlatform/third-gateway/pkg/db"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent"

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
	return nil, nil
}

func Update(ctx context.Context, in *npool.UpdateAppEmailTemplateRequest) (*npool.UpdateAppEmailTemplateResponse, error) {
	return nil, nil
}

func GetByAppLangUsedFor(ctx context.Context, in *npool.GetAppEmailTemplateByAppLangUsedForRequest) (*npool.GetAppEmailTemplateByAppLangUsedForResponse, error) {
	return nil, nil
}
