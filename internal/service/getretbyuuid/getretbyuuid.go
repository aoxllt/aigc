package getretbyuuid

import (
	v1 "aigc-go/api/getretbyuuid/v1"
	"context"
)

type Igetretbyuuid interface {
	GetRet(ctx context.Context, req *v1.GetretbyuuidReq) (string, error)
}

var localgetretbyuuid Igetretbyuuid

func GetretbyuuidforController() Igetretbyuuid {
	if localgetretbyuuid == nil {
		panic("接口为未实现")
	}
	return localgetretbyuuid
}

func RegisterGetretbyuuid(i Igetretbyuuid) {
	localgetretbyuuid = i
}
