package v1

import "github.com/gogf/gf/v2/frame/g"

type GetretReq struct {
	g.Meta     `path:"/getret" method:"get" tags:"获取结果"`
	Secret_key string `json:"secret_Key" v:"required"`
}
type GetretRes struct {
	Mes string `json:"mes"`
}
