package getretbyuuid

import (
	v1 "aigc-go/api/getretbyuuid/v1"
	"aigc-go/internal/service/getretbyuuid"
	"context"
)

func (g GetretbyuuidController) Post(ctx context.Context, req *v1.GetretbyuuidReq) (res *v1.GetretbyuuidRes, err error) {
	//TODO implement me
	res = &v1.GetretbyuuidRes{}
	res.Mes, err = getretbyuuid.GetretbyuuidforController().GetRet(ctx, req)
	if err != nil {
		return nil, err
	}
	if res.Mes == "" {
		return nil, nil
	}
	return res, nil
}
