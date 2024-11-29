package model

type PreContent struct {
	Uid           string    `json:"uid"`
	Hashuid       string    `json:"hashuid"`
	Content       string    `json:"content"`
	HandleContent string    `json:"handle_content"`
	Classes       []float64 `json:"classes"`
	Confidences   []float64 `json:"confidences"`
}

type Task struct {
	PreContent PreContent `json:"handle_content"`
	Path       []string   `json:"path"`
}

func NewTask(c PreContent, path []string) Task {
	return Task{
		PreContent: c,
		Path:       path,
	}
}

type Response struct {
	Data []string `json:"data"` // 用于存储 handlecontent 字段的数据
}
