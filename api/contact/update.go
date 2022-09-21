package contact

import (
	"context"

	tracer "github.com/NpoolPlatform/third-manager/pkg/tracer/contact"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/contact"
	contactpb "github.com/NpoolPlatform/message/npool/third/mgr/v1/contact"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	commontracer "github.com/NpoolPlatform/third-gateway/pkg/tracer"
	"github.com/NpoolPlatform/third-manager/pkg/client/contact"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateContact(ctx context.Context, in *npool.UpdateContactRequest) (*npool.UpdateContactResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateContact")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	contactInfo := &contactpb.ContactReq{
		ID:          &in.ID,
		Account:     in.Account,
		AccountType: in.AccountType,
		Sender:      in.Sender,
	}
	tracer.Trace(span, contactInfo)

	span = commontracer.TraceInvoker(span, "contact", "manager", "UpdateContact")

	info, err := contact.UpdateContact(ctx, contactInfo)
	if err != nil {
		logger.Sugar().Errorw("validate", "err", err)
		return &npool.UpdateContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateContactResponse{
		Info: info,
	}, nil
}
