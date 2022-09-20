package usedfor

import (
	"context"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/usedfor"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	usedfor.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	usedfor.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := usedfor.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
