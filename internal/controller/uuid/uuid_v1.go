package uuid

import (
	v1 "aigc-go/api/uuid/v1"
	"aigc-go/internal/service/uuid"
	"context"
)

func (u UuidController) Uuid(ctx context.Context, req *v1.UuidReq) (res *v1.UuidRes, err error) {
	res = &v1.UuidRes{}
	res.Mes, err = uuid.UuidforController().Setuuid(ctx, req)
	if err != nil {
		return res, err
	}
	return res, nil
}
