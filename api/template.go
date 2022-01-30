package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	appemailtemplatecrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appemailtemplate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppSMSTemplate(ctx context.Context, in *npool.CreateAppSMSTemplateRequest) (*npool.CreateAppSMSTemplateResponse, error) {
	return nil, nil
}

func (s *Server) GetAppSMSTemplate(ctx context.Context, in *npool.GetAppSMSTemplateRequest) (*npool.GetAppSMSTemplateResponse, error) {
	return nil, nil
}

func (s *Server) UpdateAppSMSTemplate(ctx context.Context, in *npool.UpdateAppSMSTemplateRequest) (*npool.UpdateAppSMSTemplateResponse, error) {
	return nil, nil
}

func (s *Server) GetAppSMSTemplateByAppLangUsedFor(ctx context.Context, in *npool.GetAppSMSTemplateByAppLangUsedForRequest) (*npool.GetAppSMSTemplateByAppLangUsedForResponse, error) {
	return nil, nil
}

func (s *Server) CreateAppEmailTemplate(ctx context.Context, in *npool.CreateAppEmailTemplateRequest) (*npool.CreateAppEmailTemplateResponse, error) {
	resp, err := appemailtemplatecrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app email template: %v", err)
		return &npool.CreateAppEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppEmailTemplate(ctx context.Context, in *npool.GetAppEmailTemplateRequest) (*npool.GetAppEmailTemplateResponse, error) {
	return nil, nil
}

func (s *Server) UpdateAppEmailTemplate(ctx context.Context, in *npool.UpdateAppEmailTemplateRequest) (*npool.UpdateAppEmailTemplateResponse, error) {
	return nil, nil
}

func (s *Server) GetAppEmailTemplateByAppLangUsedFor(ctx context.Context, in *npool.GetAppEmailTemplateByAppLangUsedForRequest) (*npool.GetAppEmailTemplateByAppLangUsedForResponse, error) {
	return nil, nil
}
