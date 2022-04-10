package grpc

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	appusermgrconst "github.com/NpoolPlatform/appuser-manager/pkg/message/const" //nolint
	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"

	logingwconst "github.com/NpoolPlatform/login-gateway/pkg/message/const"
	logingwpb "github.com/NpoolPlatform/message/npool/logingateway"

	"golang.org/x/xerrors"
)

const (
	grpcTimeout = 5 * time.Second
)

//---------------------------------------------------------------------------------------------------------------------------

func GetApp(ctx context.Context, in *appusermgrpb.GetAppRequest) (*appusermgrpb.GetAppResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetApp(ctx, in)
}

func GetAppUser(ctx context.Context, in *appusermgrpb.GetAppUserRequest) (*appusermgrpb.GetAppUserResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetAppUser(ctx, in)
}

func GetAppUserByAppUser(ctx context.Context, in *appusermgrpb.GetAppUserByAppUserRequest) (*appusermgrpb.GetAppUserByAppUserResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetAppUserByAppUser(ctx, in)
}

func GetAppUserSecretByAppUser(ctx context.Context, in *appusermgrpb.GetAppUserSecretByAppUserRequest) (*appusermgrpb.GetAppUserSecretByAppUserResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetAppUserSecretByAppUser(ctx, in)
}

func UpdateAppUserSecret(ctx context.Context, in *appusermgrpb.UpdateAppUserSecretRequest) (*appusermgrpb.UpdateAppUserSecretResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.UpdateAppUserSecret(ctx, in)
}

func CreateAppUserControl(ctx context.Context, in *appusermgrpb.CreateAppUserControlRequest) (*appusermgrpb.CreateAppUserControlResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.CreateAppUserControl(ctx, in)
}

func UpdateAppUserControl(ctx context.Context, in *appusermgrpb.UpdateAppUserControlRequest) (*appusermgrpb.UpdateAppUserControlResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.UpdateAppUserControl(ctx, in)
}

func GetAppUserInfoByAppUser(ctx context.Context, in *appusermgrpb.GetAppUserInfoByAppUserRequest) (*appusermgrpb.GetAppUserInfoByAppUserResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetAppUserInfoByAppUser(ctx, in)
}

// -----------------------------------------------------------------------------------------------------

func UpdateCache(ctx context.Context, in *logingwpb.UpdateCacheRequest) (*appusermgrpb.AppUserInfo, error) {
	conn, err := grpc2.GetGRPCConn(logingwconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get login gateway connection: %v", err)
	}
	defer conn.Close()

	cli := logingwpb.NewLoginGatewayClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	resp, err := cli.UpdateCache(ctx, in)
	if err != nil {
		return nil, xerrors.Errorf("fail update cache: %v", err)
	}

	return resp.Info, nil
}
