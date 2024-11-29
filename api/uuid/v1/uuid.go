package v1

import "github.com/gogf/gf/v2/frame/g"

type UuidReq struct {
	g.Meta `path:"/" method:"post" tags:"uuid设置接口" summary:"设置uuid，时长30天"`
}
type UuidRes struct {
	Mes string `json:"mes"`
}
