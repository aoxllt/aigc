package getretbyuuid

import (
	v1 "aigc-go/api/getretbyuuid/v1"
	"context"
)

type IgetrebyuuidV1 interface {
	Post(ctx context.Context, req *v1.GetretbyuuidReq) (*v1.GetretbyuuidRes, error)
}
