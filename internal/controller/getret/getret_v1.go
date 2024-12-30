package getret

import (
	v1 "aigc-go/api/getret/v1"
	"aigc-go/internal/service/getret"
	"context"
)

func (g GetretController) Get(ctx context.Context, req *v1.GetretReq) (res *v1.GetretRes, err error) {
	//TODO implement me
	res = &v1.GetretRes{}
	res.Mes, err = getret.GetretforController().Getcontent(ctx, req.Secret_key)
	if err != nil {
		return nil, err
	}
	if res.Mes == "{\"data\":[\"\"]}" {
		res.Mes = "null"
		return res, nil
	}
	return res, nil
}
