package api

import (
	"context"

	"github.com/NpoolPlatform/third-gateway/api/template/notif"

	"github.com/NpoolPlatform/third-gateway/api/template/sms"
	"github.com/NpoolPlatform/third-gateway/api/verify"

	v1 "github.com/NpoolPlatform/message/npool/third/gw/v1"
	"github.com/NpoolPlatform/third-gateway/api/contact"
	"github.com/NpoolPlatform/third-gateway/api/template/email"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterGatewayServer(server, &Server{})
	contact.Register(server)
	email.Register(server)
	sms.Register(server)
	notif.Register(server)
	verify.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := contact.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := email.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := sms.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := verify.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
