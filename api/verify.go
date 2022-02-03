package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	emailmw "github.com/NpoolPlatform/third-gateway/pkg/middleware/email"
	googlemw "github.com/NpoolPlatform/third-gateway/pkg/middleware/google"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendSMSCode(ctx context.Context, in *npool.SendSMSCodeRequest) (*npool.SendSMSCodeResponse, error) {
	return nil, nil
}

func (s *Server) VerifySMSCode(ctx context.Context, in *npool.VerifySMSCodeRequest) (*npool.VerifySMSCodeResponse, error) {
	return nil, nil
}

func (s *Server) SendEmailCode(ctx context.Context, in *npool.SendEmailCodeRequest) (*npool.SendEmailCodeResponse, error) {
	resp, err := emailmw.SendCode(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail send email code: %v", err)
		return &npool.SendEmailCodeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) VerifyEmailCode(ctx context.Context, in *npool.VerifyEmailCodeRequest) (*npool.VerifyEmailCodeResponse, error) {
	resp, err := emailmw.VerifyCode(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail verify email code: %v", err)
		return &npool.VerifyEmailCodeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) SetupGoogleAuthentication(ctx context.Context, in *npool.SetupGoogleAuthenticationRequest) (*npool.SetupGoogleAuthenticationResponse, error) {
	resp, err := googlemw.SetupGoogleAuthentication(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail setup google authentication: %v", err)
		return &npool.SetupGoogleAuthenticationResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) VerifyGoogleAuthentication(ctx context.Context, in *npool.VerifyGoogleAuthenticationRequest) (*npool.VerifyGoogleAuthenticationResponse, error) {
	resp, err := googlemw.VerifyGoogleAuthentication(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail verify google authentication: %v", err)
		return &npool.VerifyGoogleAuthenticationResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
