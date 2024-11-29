package upload

import (
	v1 "aigc-go/api/upload/v1"
	"context"
)

type IuploadV1 interface {
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
}
