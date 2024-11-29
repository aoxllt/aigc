package getretbyuuid

import "aigc-go/api/getretbyuuid"

type GetretbyuuidController struct{}

func NewGetretbyuuidV1() getretbyuuid.IgetrebyuuidV1 {
	return &GetretbyuuidController{}
}
