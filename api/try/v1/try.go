package v1

import "github.com/gogf/gf/v2/frame/g"

type TryReq struct {
	g.Meta `path:"/try" method:"post" tags:"尝试接口" summary:"上传200字内的语句"`
	Text   string `json:"text" v:"required"`
}
type TryRes struct {
	Mes string `json:"mes"`
}
