package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	emailmw "github.com/NpoolPlatform/third-gateway/pkg/middleware/email"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) NotifyEmail(ctx context.Context, in *npool.NotifyEmailRequest) (*npool.NotifyEmailResponse, error) {
	notif, err := emailmw.Notify(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail notify email: %v", err)
		return &npool.NotifyEmailResponse{}, status.Error(codes.Internal, err.Error())
	}
	return notif, nil
}
