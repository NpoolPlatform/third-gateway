package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
)

func (s *Server) SendSMSCode(ctx context.Context, in *npool.SendSMSCodeRequest) (*npool.SendSMSCodeResponse, error) {
	return nil, nil
}

func (s *Server) VerifySMSCode(ctx context.Context, in *npool.VerifySMSCodeRequest) (*npool.VerifySMSCodeResponse, error) {
	return nil, nil
}

func (s *Server) SendEmailCode(ctx context.Context, in *npool.SendEmailCodeRequest) (*npool.SendEmailCodeResponse, error) {
	return nil, nil
}

func (s *Server) VerifyEmailCode(ctx context.Context, in *npool.VerifyEmailCodeRequest) (*npool.VerifyEmailCodeResponse, error) {
	return nil, nil
}
