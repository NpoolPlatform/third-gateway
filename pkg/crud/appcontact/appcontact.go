package appcontact

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"

	"github.com/NpoolPlatform/third-gateway/pkg/db"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appcontact"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateTemplate(info *npool.AppContact) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if info.GetAccount() == "" {
		return xerrors.Errorf("invalid account")
	}
	return nil
}

func dbRowToTemplate(row *ent.AppContact) *npool.AppContact {
	return &npool.AppContact{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		UsedFor:     row.UsedFor,
		Account:     row.Account,
		AccountType: row.AccountType,
		Sender:      row.Sender,
	}
}

func Create(ctx context.Context, in *npool.CreateAppContactRequest) (*npool.CreateAppContactResponse, error) {
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
		AppContact.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUsedFor(in.GetInfo().GetUsedFor()).
		SetAccount(in.GetInfo().GetAccount()).
		SetAccountType(in.GetInfo().GetAccountType()).
		SetSender(in.GetInfo().GetSender()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app contact: %v", err)
	}

	return &npool.CreateAppContactResponse{
		Info: dbRowToTemplate(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppContactRequest) (*npool.GetAppContactResponse, error) {
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
		AppContact.
		Query().
		Where(
			appcontact.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app contact: %v", err)
	}

	var template *npool.AppContact
	for _, info := range infos {
		template = dbRowToTemplate(info)
		break
	}

	return &npool.GetAppContactResponse{
		Info: template,
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppContactRequest) (*npool.UpdateAppContactResponse, error) {
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
		AppContact.
		UpdateOneID(id).
		SetAccount(in.GetInfo().GetAccount()).
		SetAccountType(in.GetInfo().GetAccountType()).
		SetSender(in.GetInfo().GetSender()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app contact: %v", err)
	}

	return &npool.UpdateAppContactResponse{
		Info: dbRowToTemplate(info),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppContactsByAppRequest) (*npool.GetAppContactsByAppResponse, error) {
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
		AppContact.
		Query().
		Where(
			appcontact.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app contact: %v", err)
	}

	templates := []*npool.AppContact{}
	for _, info := range infos {
		templates = append(templates, dbRowToTemplate(info))
	}

	return &npool.GetAppContactsByAppResponse{
		Infos: templates,
	}, nil
}

func GetByAppUsedForAccountType(ctx context.Context, in *npool.GetAppContactByAppUsedForAccountTypeRequest) (*npool.GetAppContactByAppUsedForAccountTypeResponse, error) {
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
		AppContact.
		Query().
		Where(
			appcontact.And(
				appcontact.AppID(appID),
				appcontact.UsedFor(in.GetUsedFor()),
				appcontact.AccountType(in.GetAccountType()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app contact: %v", err)
	}

	var template *npool.AppContact
	for _, info := range infos {
		template = dbRowToTemplate(info)
		break
	}

	return &npool.GetAppContactByAppUsedForAccountTypeResponse{
		Info: template,
	}, nil
}
