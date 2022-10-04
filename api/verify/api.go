package verify

import (
	"context"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/verify"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	verify.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	verify.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := verify.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
