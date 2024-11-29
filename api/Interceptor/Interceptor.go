package Interceptor

import (
	v1 "aigc-go/api/Interceptor/v1"
	"context"
)

type IinterceptorV1 interface {
	Post(ctx context.Context, req *v1.InterceptorReq) (*v1.InterceptorRes, error)
}
