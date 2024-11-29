package getret

import (
	v1 "aigc-go/api/getret/v1"
	"context"
)

type IGetRetV1 interface {
	Get(ctx context.Context, req *v1.GetretReq) (*v1.GetretRes, error)
}
