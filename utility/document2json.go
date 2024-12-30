package utility

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"os"
	"strings"
)

func init() {

}
func Document2Json(filepath string) (map[string][]string, error) {
	var pyt Python
	content := pyt.Handle("." + filepath)
	data, err := ProcessTextToJSON(content, 200)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ProcessTextToJSON(text string, maxLength int) (map[string][]string, error) {
	splitParagraphs := func(text string) []string {
		paragraphs := strings.Split(text, "\n")
		var result []string
		for _, p := range paragraphs {
			p = strings.TrimSpace(p)
			if p != "" {
				result = append(result, p)
			}
		}
		return result
	}

	// splitLongParagraph 将超过 maxLength 的段落分割成多个部分
	splitLongParagraph := func(paragraph string, maxLength int) []string {
		if len(paragraph) <= maxLength {
			return []string{paragraph}
		} else {
			var parts []string
			for len(paragraph) > maxLength {
				splitIndex := maxLength
				// 寻找最后一个完整句子的分隔符
				for splitIndex > 0 && !strings.ContainsAny(string(paragraph[splitIndex]), "。！？.?!") {
					splitIndex--
				}
				if splitIndex == 0 {
					splitIndex = maxLength
				} else {
					splitIndex++
				}

				// 截取并保存段落
				parts = append(parts, strings.TrimSpace(paragraph[:splitIndex]))
				paragraph = strings.TrimSpace(paragraph[splitIndex:])
			}

			// 最后剩余的段落
			parts = append(parts, paragraph)
			return parts
		}
	}

	// 获取所有的段落
	paragraphs := splitParagraphs(text)

	// 用来存储结果数据，直接使用 map
	result := make(map[string][]string)
	for i, para := range paragraphs {
		// 对每个段落进行处理，分割成较小的部分
		parts := splitLongParagraph(para, maxLength)
		// 按照段落顺序存储，key 使用 paragraph1, paragraph2 等
		result[fmt.Sprintf("paragraph%d", i+1)] = parts
	}

	return result, nil
}

func Next(result map[string][]string) (string, error) {
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
func SaveFile(file *ghttp.UploadFile) (string, error) {
	tmpFilePath := "./static/tmp_" + file.Filename
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		return "", fmt.Errorf("无法创建临时文件: %w", err)
	}
	defer tmpFile.Close()

	// 打开上传的文件
	fileReader, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("无法打开上传文件: %w", err)
	}
	defer fileReader.Close()

	// 将文件内容写入临时文件
	_, err = io.Copy(tmpFile, fileReader)
	if err != nil {
		return "", fmt.Errorf("无法将文件内容写入临时文件: %w", err)
	}

	return "\\static\\tmp_" + file.Filename, nil
}
