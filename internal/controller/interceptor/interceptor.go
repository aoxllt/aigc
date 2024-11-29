package interceptor

import "aigc-go/api/Interceptor"

type InterceptorController struct{}

func NewInterceptorV1() Interceptor.IinterceptorV1 {
	return &InterceptorController{}
}
