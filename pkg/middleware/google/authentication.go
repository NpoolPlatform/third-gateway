package google

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"
	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	grpc2 "github.com/NpoolPlatform/third-gateway/pkg/grpc"

	"golang.org/x/xerrors"
)

func SetupGoogleAuthentication(ctx context.Context, in *npool.SetupGoogleAuthenticationRequest) (*npool.SetupGoogleAuthenticationResponse, error) {
	resp, err := grpc2.GetAppUserInfoByAppUser(ctx, &appusermgrpb.GetAppUserInfoByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("invalid app user: %v", err)
	}
	if resp.Info == nil {
		return nil, xerrors.Errorf("invalid app user")
	}

	resp1, err := grpc2.GetAppUserSecretByAppUser(ctx, &appusermgrpb.GetAppUserSecretByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("invalid app user secret: %v", err)
	}
	if resp1.Info == nil {
		return nil, xerrors.Errorf("invalid app user secret")
	}

	generateSecret := false
	if resp.Info.Ctrl == nil || !resp.Info.Ctrl.GoogleAuthenticationVerified {
		generateSecret = true
	}
	if resp1.Info.GoogleSecret == "" {
		generateSecret = true
	}

	if generateSecret {
		secret, err := GenerateSecret()
		if err != nil {
			return nil, xerrors.Errorf("fail generate google secret: %v", err)
		}
		resp1.Info.GoogleSecret = secret
		_, err = grpc2.UpdateAppUserSecret(ctx, &appusermgrpb.UpdateAppUserSecretRequest{
			Info: resp1.Info,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail update google secret: %v", err)
		}
	}

	account := resp.Info.User.EmailAddress
	if account == "" {
		account = resp.Info.User.PhoneNO
	}
	if account == "" {
		return nil, xerrors.Errorf("invalid user account info")
	}

	return &npool.SetupGoogleAuthenticationResponse{
		OTPAuth: fmt.Sprintf("otpauth://totp/%s?secret=%s", account, resp1.Info.GoogleSecret),
		Secret:  resp1.Info.GoogleSecret,
	}, nil
}

func VerifyGoogleAuthentication(ctx context.Context, in *npool.VerifyGoogleAuthenticationRequest) (*npool.VerifyGoogleAuthenticationResponse, error) {
	resp, err := grpc2.GetAppUserInfoByAppUser(ctx, &appusermgrpb.GetAppUserInfoByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("invalid app user: %v", err)
	}
	if resp.Info == nil {
		return nil, xerrors.Errorf("invalid app user")
	}

	resp1, err := grpc2.GetAppUserSecretByAppUser(ctx, &appusermgrpb.GetAppUserSecretByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("invalid app user secret: %v", err)
	}
	if resp1.Info == nil {
		return nil, xerrors.Errorf("invalid app user secret")
	}

	ok, err := VerifyCode(resp1.Info.GoogleSecret, in.GetCode())
	if err != nil {
		return nil, xerrors.Errorf("fail verify google code: %v", err)
	}

	if ok {
		if resp.Info.Ctrl == nil {
			_, err = grpc2.CreateAppUserControl(ctx, &appusermgrpb.CreateAppUserControlRequest{
				Info: &appusermgrpb.AppUserControl{
					AppID:                        in.GetAppID(),
					UserID:                       in.GetUserID(),
					GoogleAuthenticationVerified: true,
				},
			})
		} else {
			resp.Info.Ctrl.GoogleAuthenticationVerified = true
			_, err = grpc2.UpdateAppUserControl(ctx, &appusermgrpb.UpdateAppUserControlRequest{
				Info: resp.Info.Ctrl,
			})
		}
		if err != nil {
			return nil, xerrors.Errorf("fail set app user control: %v", err)
		}
	}

	logger.Sugar().Infof("verify google secret %v code %v: %v", resp1.Info.GoogleSecret, in.GetCode(), ok)

	var code int32
	if !ok {
		code = -1
	}

	return &npool.VerifyGoogleAuthenticationResponse{
		Code: code,
	}, nil
}
