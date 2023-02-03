package notif

import (
	"context"

	"github.com/NpoolPlatform/message/npool/third/gw/v1/template/notif"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	notif.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	notif.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := notif.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
