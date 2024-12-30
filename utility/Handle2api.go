package utility

import (
	"aigc-go/internal/consts"
	"aigc-go/internal/dao"
	"aigc-go/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"math"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type APIResponse struct {
	Classes     []float64 `json:"classes"`
	Confidences []float64 `json:"confidences"`
}
type FinalResult struct {
	Data []map[string]interface{} `json:"data"`
}

func init() {

}

var TaskChannel = make(chan model.Task, 10) // 缓冲区大小根据需求调整

// 启动 Handle2api 处理任务
func Handle2api() {
	go func() {

		for t := range TaskChannel {
			// 遍历任务中的文件
			for _, filepath := range t.Path {
				// 文件转换为 JSON
				data, err := Document2Json(filepath)
				if err != nil {
					fmt.Printf("文件转换为 JSON 失败: %v\n", err)
					continue
				}

				json, err := Next(data)
				if err != nil {
					return
				}
				if err != nil {
					return
				}
				// 更新任务内容
				t.PreContent.Content = json
				t.PreContent.HandleContent = ""

				// 使用独立上下文处理数据库操作
				dbCtx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
				defer cancel()

				// 插入数据库
				_, err = dao.Content.Ctx(dbCtx).Insert(t.PreContent)
				if err != nil {
					fmt.Printf("数据库插入失败: %v\n", err)
					continue
				}

				// 调用 API 处理内容
				retContent, err := UploadToAPI(t.PreContent.Content)
				if err != nil {
					fmt.Printf("处理 API 请求失败: %v\n", err)
					continue
				}
				handleContent, err := ProcessResponseData(retContent, data, filepath)
				fmt.Println(handleContent)
				if err != nil {
					return
				}
				// 更新数据库
				_, err = dao.Content.Ctx(dbCtx).Where("hashuid=? and content=?", t.PreContent.Hashuid, t.PreContent.Content).Data(g.Map{
					"HandleContent": handleContent,
					"handletime":    time.Now(),
				}).Update()
				if err != nil {
					fmt.Printf("更新数据库失败: %v\n", err)
				}
			}
		}
	}()

}

func ProcessResponseData(response *http.Response, data map[string][]string, path string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	path_default := "try" + strconv.Itoa(rand.Intn(10000))
	if len(path) == 0 {
		path = path_default
	}
	defer response.Body.Close()
	var apiResponse map[string]APIResponse
	err := json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Errorf("解析响应数据失败: %v", err)
		return " ", nil
	}
	averagedData := make(map[string]APIResponse)
	for key, value := range apiResponse {
		averageClass := average(value.Classes)
		averageConfidence := average(value.Confidences)

		// 更新数据为平均值
		averagedData[key] = APIResponse{
			Classes:     []float64{averageClass},
			Confidences: []float64{averageConfidence},
		}
	}
	totalWords := 0

	// 遍历 data map
	for key, value := range data {
		// 如果值是字符串切片 (list)
		if len(value) > 0 {
			// 合并切片中的字符串
			combinedText := strings.Join(value, "")
			// 去除所有空格
			combinedTextNoSpaces := strings.ReplaceAll(combinedText, " ", "")
			// 更新 map 中的值为合并后的文本
			data[key] = []string{combinedText}
			// 计算总字数（去掉空格后的字符数）
			totalWords += len(combinedTextNoSpaces)
		}
	}
	_, fileName := filepath.Split(path)

	// 去掉 "tmp_" 前缀
	fileName = strings.TrimPrefix(fileName, "tmp_")
	result := createFinalJSON(data, averagedData, fileName, totalWords)
	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	return string(resultJSON), nil
}
func average(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func UploadToAPI(jsonData string) (*http.Response, error) {
	// 将字符串解析为 JSON 格式
	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, fmt.Errorf("解析字符串为 JSON 时出错: %v", err)
	}

	// 将 JSON 数据重新编码为字节流
	requestBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("将 JSON 数据编码为字节流时出错: %v", err)
	}

	// 如果没有特别需要去除首尾字符，直接发送请求
	resp, err := http.Post(consts.ApiUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("向 API 上传数据时出错: %v", err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API 请求失败，状态码：%d", resp.StatusCode)
	}
	return resp, nil

}

func createFinalJSON(jsonData1 map[string][]string, jsonData2 map[string]APIResponse, fileName string, totalWords int) map[string]interface{} {
	result := FinalResult{
		Data: []map[string]interface{}{},
	}

	// 初始化内容变量
	content := []map[string]interface{}{}
	totalSuspectedText := 0
	totalConfidence := 0.0
	paragraphCount := 0

	// 遍历 jsonData1
	for key, value1 := range jsonData1 {
		// 跳过 title
		if key == "title" {
			continue
		}

		// 检查 jsonData2 是否有该 key
		if value2, exists := jsonData2[key]; exists {
			// 合并文本
			paragraphText := strings.Join(value1, "")
			numWords := len(strings.ReplaceAll(paragraphText, " ", ""))

			// 计算 suspectedText
			classes := value2.Classes[0]
			suspectedText := math.Floor(float64(numWords) * classes)
			proportion := round(suspectedText/float64(totalWords), 2)
			confidence := value2.Confidences[0]

			// 更新统计数据
			paragraphCount++
			totalSuspectedText += int(suspectedText)
			totalConfidence += confidence

			// 将内容添加到 content 列表
			content = append(content, map[string]interface{}{
				key:             paragraphText,
				"classes":       classes,
				"suspectedtext": int(suspectedText),
				"proportion":    proportion,
				"confidence":    confidence,
				"words":         numWords,
			})
		}
	}

	// 计算平均比例和平均置信度
	averProportion := round(float64(totalSuspectedText)/float64(totalWords), 2)
	aveConfidence := round(totalConfidence/float64(paragraphCount), 2)

	// 将最终数据添加到 result
	result.Data = append(result.Data, map[string]interface{}{
		"title":         fileName,
		"totalwords":    fmt.Sprintf("%d", totalWords),
		"aveproportion": averProportion,
		"aveconfidence": aveConfidence,
		"content":       content,
	})

	// 返回结果
	return result.Data[0]
}

func round(val float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Round(val*shift) / shift
}
