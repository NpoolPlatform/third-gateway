package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/thirdgateway"
	appcontactcrud "github.com/NpoolPlatform/third-gateway/pkg/crud/appcontact"
	emailmw "github.com/NpoolPlatform/third-gateway/pkg/middleware/email"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ContactByEmail(ctx context.Context, in *npool.ContactByEmailRequest) (*npool.ContactByEmailResponse, error) {
	resp, err := emailmw.Contact(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail contact by email: %v", err)
		return &npool.ContactByEmailResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppContact(ctx context.Context, in *npool.CreateAppContactRequest) (*npool.CreateAppContactResponse, error) {
	resp, err := appcontactcrud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create app contact: %v", err)
		return &npool.CreateAppContactResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppContact(ctx context.Context, in *npool.GetAppContactRequest) (*npool.GetAppContactResponse, error) {
	resp, err := appcontactcrud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app contact: %v", err)
		return &npool.GetAppContactResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppContact(ctx context.Context, in *npool.UpdateAppContactRequest) (*npool.UpdateAppContactResponse, error) {
	resp, err := appcontactcrud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update app contact: %v", err)
		return &npool.UpdateAppContactResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppContactsByApp(ctx context.Context, in *npool.GetAppContactsByAppRequest) (*npool.GetAppContactsByAppResponse, error) {
	resp, err := appcontactcrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app contacts by app: %v", err)
		return &npool.GetAppContactsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppContactsByOtherApp(ctx context.Context, in *npool.GetAppContactsByOtherAppRequest) (*npool.GetAppContactsByOtherAppResponse, error) {
	resp, err := appcontactcrud.GetByApp(ctx, &npool.GetAppContactsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app contacts by other app: %v", err)
		return &npool.GetAppContactsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppContactsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetAppContactByAppUsedForAccountType(ctx context.Context, in *npool.GetAppContactByAppUsedForAccountTypeRequest) (*npool.GetAppContactByAppUsedForAccountTypeResponse, error) {
	resp, err := appcontactcrud.GetByAppUsedForAccountType(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app contact by app used for account type: %v", err)
		return &npool.GetAppContactByAppUsedForAccountTypeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
