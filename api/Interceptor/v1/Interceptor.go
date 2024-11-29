package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type InterceptorReq struct {
	g.Meta `path:"/interceptor" method:"POST" tags:"检测器" summary:"检测是否设置uuid"`
}
type InterceptorRes struct {
	Mes string `json:"mes"`
}
