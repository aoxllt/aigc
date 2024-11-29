package v1

import "github.com/gogf/gf/v2/frame/g"

type GetretbyuuidReq struct {
	g.Meta `path:"/getretbyuuid" method:"POST" tags:"通过uuid获取处理的内容"`
}
type GetretbyuuidRes struct {
	Mes string `json:"mes"`
}
