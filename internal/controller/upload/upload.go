package upload

import (
	"aigc-go/api/upload"
)

type UploadController struct{}

func Uploadv1() upload.IuploadV1 {
	return &UploadController{}
}
