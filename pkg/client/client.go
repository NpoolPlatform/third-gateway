package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
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
