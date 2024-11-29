package upload

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Iupload interface {
	Handle(context.Context, []*ghttp.UploadFile) (string, error)
}

var localUploadservice Iupload

func UploadforController() Iupload {
	if localUploadservice == nil {
		panic("localUploadservice is nil")
	}
	return localUploadservice
}
func RegisterUpolad(c Iupload) {
	localUploadservice = c
}
