package getret

import (
	"aigc-go/internal/dao"
	"aigc-go/internal/service/getret"
	"context"
	"encoding/json"
	"fmt"
)

func init() {
	getret.Registergetret(GetRet{})
}

type GetRet struct{}
type Response struct {
	Data []string `json:"data"` // 用于存储 handlecontent 字段的数据
}

func (g GetRet) Getcontent(ctx context.Context, sk string) (string, error) {
	//TODO implement me
	//拦截器
	all, err := dao.Content.Ctx(ctx).Where("hashuid = ?", sk).All()
	if err != nil {
		// 查询失败，返回错误
		return "", fmt.Errorf("查询失败: %v", err)
	}

	// 创建一个切片用于存储 handlecontent 字段的值
	var contents []string

	// 遍历查询结果并提取 handlecontent 字段
	for _, record := range all {
		// 获取 handlecontent 字段
		if handlecontent, exists := record["handlecontent"]; exists {
			contents = append(contents, handlecontent.String())
		}
	}
	res := Response{Data: contents}

	// 将 Response 结构体序列化为 JSON 字符串
	jsonData, err := json.Marshal(res)
	if err != nil {
		return "", fmt.Errorf("序列化 JSON 失败: %v", err)
	}

	// 返回 JSON 字符串
	return string(jsonData), nil
}
