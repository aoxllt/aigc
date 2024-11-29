package upload

import (
	"aigc-go/internal/model"
	"aigc-go/internal/service/upload"
	"aigc-go/utility"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

func init() {
	upload.RegisterUpolad(uploadreg{})
}

type uploadreg struct{}

func (u uploadreg) Handle(ctx context.Context, files []*ghttp.UploadFile) (string, error) {
	var c model.PreContent
	var task model.Task
	path := make([]string, len(files))

	// 获取用户 UUID
	get, err := ghttp.RequestFromCtx(ctx).Session.Get("uuid")
	if err != nil {
		return "false", err
	}
	if get.String() == "" {
		return "false,uuid为空", nil
	}
	uuid := get.String()
	randomSeed := fmt.Sprintf("%s-%d", uuid, time.Now().UnixNano())
	hash := sha1.Sum([]byte(randomSeed))
	hashedUUID := hex.EncodeToString(hash[:])

	c.Uid = uuid
	c.Hashuid = hashedUUID

	for i, file := range files {
		if file.Filename == "" {
			return "存在空文件名", nil
		}
		path[i], err = utility.SaveFile(file)
		if err != nil {
			return "false", err
		}
	}
	task = model.NewTask(c, path)
	utility.TaskChannel <- task

	// 返回 hashuid
	return hashedUUID, nil
}
