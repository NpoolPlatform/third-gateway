package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	appemailtemplatecrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appemailtemplate"
	appsmstemplatecrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appsmstemplate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppSMSTemplate(ctx context.Context, in *npool.CreateAppSMSTemplateRequest) (*npool.CreateAppSMSTemplateResponse, error) {
	resp, err := appsmstemplatecrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app sms template: %v", err)
		return &npool.CreateAppSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppSMSTemplate(ctx context.Context, in *npool.GetAppSMSTemplateRequest) (*npool.GetAppSMSTemplateResponse, error) {
	resp, err := appsmstemplatecrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app sms template: %v", err)
		return &npool.GetAppSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppSMSTemplate(ctx context.Context, in *npool.UpdateAppSMSTemplateRequest) (*npool.UpdateAppSMSTemplateResponse, error) {
	resp, err := appsmstemplatecrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app sms template: %v", err)
		return &npool.UpdateAppSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppSMSTemplatesByApp(ctx context.Context, in *npool.GetAppSMSTemplatesByAppRequest) (*npool.GetAppSMSTemplatesByAppResponse, error) {
	resp, err := appsmstemplatecrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app sms templates by app: %v", err)
		return &npool.GetAppSMSTemplatesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppSMSTemplatesByOtherApp(ctx context.Context, in *npool.GetAppSMSTemplatesByOtherAppRequest) (*npool.GetAppSMSTemplatesByOtherAppResponse, error) {
	resp, err := appsmstemplatecrud.GetByApp(ctx, &npool.GetAppSMSTemplatesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app sms templates by other app: %v", err)
		return &npool.GetAppSMSTemplatesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppSMSTemplatesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetAppSMSTemplateByAppLangUsedFor(ctx context.Context, in *npool.GetAppSMSTemplateByAppLangUsedForRequest) (*npool.GetAppSMSTemplateByAppLangUsedForResponse, error) {
	resp, err := appsmstemplatecrud.GetByAppLangUsedFor(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app sms template by app lang used for: %v", err)
		return &npool.GetAppSMSTemplateByAppLangUsedForResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
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
	resp, err := appemailtemplatecrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app email template: %v", err)
		return &npool.GetAppEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppEmailTemplate(ctx context.Context, in *npool.UpdateAppEmailTemplateRequest) (*npool.UpdateAppEmailTemplateResponse, error) {
	resp, err := appemailtemplatecrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app email template: %v", err)
		return &npool.UpdateAppEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppEmailTemplatesByApp(ctx context.Context, in *npool.GetAppEmailTemplatesByAppRequest) (*npool.GetAppEmailTemplatesByAppResponse, error) {
	resp, err := appemailtemplatecrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app email templates by app: %v", err)
		return &npool.GetAppEmailTemplatesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppEmailTemplatesByOtherApp(ctx context.Context, in *npool.GetAppEmailTemplatesByOtherAppRequest) (*npool.GetAppEmailTemplatesByOtherAppResponse, error) {
	resp, err := appemailtemplatecrud.GetByApp(ctx, &npool.GetAppEmailTemplatesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app email templates by other app: %v", err)
		return &npool.GetAppEmailTemplatesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppEmailTemplatesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetAppEmailTemplateByAppLangUsedFor(ctx context.Context, in *npool.GetAppEmailTemplateByAppLangUsedForRequest) (*npool.GetAppEmailTemplateByAppLangUsedForResponse, error) {
	resp, err := appemailtemplatecrud.GetByAppLangUsedFor(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app email template by app lang used for: %v", err)
		return &npool.GetAppEmailTemplateByAppLangUsedForResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
