package try

import (
	try2 "aigc-go/internal/service/try"
	"aigc-go/utility"
	"context"
)

func init() {
	try2.RegisterItry(try{})
}

type RequestBody struct {
	Paragraph1 []string `json:"paragraph1"`
}

type try struct{}

func (t try) Process(ctx context.Context, text string) (string, error) {
	//TODO implement me
	json, err := utility.ProcessTextToJSON(text, 200)
	if err != nil {
		return "", err
	}
	next, err := utility.Next(json)
	if err != nil {
		return "", err
	}
	retContent, err := utility.UploadToAPI(next)
	if err != nil {
		return "", err
	}
	data, err := utility.ProcessResponseData(retContent, json, "")
	if err != nil {
		return "", err
	}
	return data, nil
}
