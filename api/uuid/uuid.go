package uuid

import (
	v1 "aigc-go/api/uuid/v1"
	"context"
)

type Iuuidv1 interface {
	Uuid(ctx context.Context, req *v1.UuidReq) (*v1.UuidRes, error)
}
