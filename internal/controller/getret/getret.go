package getret

import "aigc-go/api/getret"

type GetretController struct {
}

func NewGetretV1() getret.IGetRetV1 {
	return &GetretController{}
}
