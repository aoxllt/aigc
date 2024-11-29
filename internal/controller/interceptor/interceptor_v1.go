package interceptor

import (
	v1 "aigc-go/api/Interceptor/v1"
	"aigc-go/internal/service/uuid"
	"context"
)

func (i InterceptorController) Post(ctx context.Context, req *v1.InterceptorReq) (*v1.InterceptorRes, error) {
	//TODO implement me
	res := &v1.InterceptorRes{}
	getuuid, err := uuid.UuidforController().Getuuid(ctx)
	if err != nil {
		return nil, err
	}
	if getuuid == "" {
		res.Mes = "false"
		return res, nil
	}
	res.Mes = "true"
	return res, nil
}
