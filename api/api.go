package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedThirdGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterThirdGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterThirdGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
