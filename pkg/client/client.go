package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"

	signmethodpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.ThirdGatewayClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get third gateway connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewThirdGatewayClient(conn)

	return fn(_ctx, cli)
}

func NotifyEmail(ctx context.Context, in *npool.NotifyEmailRequest) error {
	info, err := do(ctx, func(_ctx context.Context, cli npool.ThirdGatewayClient) (cruder.Any, error) {
		resp, err := cli.NotifyEmail(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail notify email: %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return fmt.Errorf("fail notify email: %v", err)
	}

	if info.(*npool.NotifyEmailResponse).Code < 0 {
		return fmt.Errorf("fail notify email: %v", err)
	}

	return nil
}

func VerifySMSCode(ctx context.Context, in *npool.VerifySMSCodeRequest) error {
	info, err := do(ctx, func(_ctx context.Context, cli npool.ThirdGatewayClient) (cruder.Any, error) {
		resp, err := cli.VerifySMSCode(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return err
	}

	if info.(*npool.VerifySMSCodeResponse).Code < 0 {
		return fmt.Errorf("fail verify SMS code")
	}
	return nil
}

func VerifyEmailCode(ctx context.Context, in *npool.VerifyEmailCodeRequest) error {
	info, err := do(ctx, func(_ctx context.Context, cli npool.ThirdGatewayClient) (cruder.Any, error) {
		resp, err := cli.VerifyEmailCode(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return err
	}

	if info.(*npool.VerifyEmailCodeResponse).Code < 0 {
		return fmt.Errorf("fail verify email code")
	}

	return nil
}

func VerifyGoogleAuthentication(ctx context.Context, in *npool.VerifyGoogleAuthenticationRequest) error {
	info, err := do(ctx, func(_ctx context.Context, cli npool.ThirdGatewayClient) (cruder.Any, error) {
		resp, err := cli.VerifyGoogleAuthentication(ctx, in)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return err
	}

	if info.(*npool.VerifyGoogleAuthenticationResponse).Code < 0 {
		return fmt.Errorf("fail verify google authentication")
	}

	return nil
}

func VerifyCode(ctx context.Context, appID, userID string, signMethod signmethodpb.SignMethodType, account, code, usedFor string) error {
	switch signMethod {
	case signmethodpb.SignMethodType_Email:
		return VerifyEmailCode(ctx, &npool.VerifyEmailCodeRequest{
			AppID:        appID,
			UserID:       userID,
			EmailAddress: account,
			UsedFor:      usedFor,
			Code:         code,
		})
	case signmethodpb.SignMethodType_Mobile:
		return VerifySMSCode(ctx, &npool.VerifySMSCodeRequest{
			AppID:   appID,
			UserID:  userID,
			PhoneNO: account,
			UsedFor: usedFor,
			Code:    code,
		})
	case signmethodpb.SignMethodType_Google:
		return VerifyGoogleAuthentication(ctx, &npool.VerifyGoogleAuthenticationRequest{
			AppID:  appID,
			UserID: userID,
			Code:   code,
		})
	}
	return fmt.Errorf("unknown sign method")
}

func VerifyGoogleRecaptchaV3(ctx context.Context, recaptchaToken string) error {
	_, err := do(ctx, func(_ctx context.Context, cli npool.ThirdGatewayClient) (cruder.Any, error) {
		_, err := cli.VerifyGoogleRecaptchaV3(ctx, &npool.VerifyGoogleRecaptchaV3Request{
			RecaptchaToken: recaptchaToken,
		})
		return nil, err
	})
	return err
}
