package usedfor

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/third/gw/v1/usedfor"
	constant "github.com/NpoolPlatform/third-gateway/pkg/message/const"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
)

func (s *Server) GetUsedFor(ctx context.Context, in *npool.GetUsedForRequest) (*npool.GetUsedForResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateContact")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	infos := []usedfor.UsedFor{}

	for key := range usedfor.UsedFor_value {
		usedFor := usedfor.UsedFor(usedfor.UsedFor_value[key]).Enum()
		infos = append(infos, *usedFor)
	}

	return &npool.GetUsedForResponse{
		UsedFor: infos,
	}, nil
}
