package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta `path:"/upload" method:"POST" tags:"文件上传" summary:"上传待处理的文件"`
	Files  []*ghttp.UploadFile `json:"files" v:"required"`
}

type UploadRes struct {
	Mes string `json:"mes"`
}
