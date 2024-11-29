package upload

import (
	"aigc-go/api/upload/v1"
	"aigc-go/internal/service/upload"
	"context"
)

func (c *UploadController) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	res = &v1.UploadRes{}
	res.Mes, err = upload.UploadforController().Handle(ctx, req.Files)
	if err != nil {
		return nil, err
	}
	if res.Mes == "false" {
		res.Mes = "服务器出错"
		return res, nil
	}
	return res, nil
}
