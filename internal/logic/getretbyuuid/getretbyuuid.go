package getretbyuuid

import (
	v1 "aigc-go/api/getretbyuuid/v1"
	"aigc-go/internal/dao"
	"aigc-go/internal/model"
	"aigc-go/internal/service/getretbyuuid"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)
import uid "aigc-go/internal/service/uuid"

func init() {
	getretbyuuid.RegisterGetretbyuuid(&Getretbyuuid{})
}

type Getretbyuuid struct{}

func (g Getretbyuuid) GetRet(ctx context.Context, req *v1.GetretbyuuidReq) (string, error) {
	//TODO implement me
	getuuid, err := uid.UuidforController().Getuuid(ctx)
	if err != nil {
		return "", err
	}
	if getuuid == "" {
		return "", errors.New("uuid empty")
	}
	all, err := dao.Content.Ctx(ctx).Where("uid=?", getuuid).All()
	if err != nil {
		return "", err
	}
	var contents []string

	// 遍历查询结果并提取 handlecontent 字段
	for _, record := range all {
		// 获取 handlecontent 字段
		if handlecontent, exists := record["handlecontent"]; exists {
			contents = append(contents, handlecontent.String(), record["handletime"].String())
		}
	}
	res := model.Response{Data: contents}

	// 将 Response 结构体序列化为 JSON 字符串
	jsonData, err := json.Marshal(res)
	if err != nil {
		return "", fmt.Errorf("序列化 JSON 失败: %v", err)
	}

	// 返回 JSON 字符串
	return string(jsonData), nil
}
