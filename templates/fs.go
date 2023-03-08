package templates

import (
	"embed"
	"encoding/json"
)

//go:embed *
var FS embed.FS

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func GetFAQs() []Question {
	data, _ := FS.ReadFile("faq.json")
	var questions []Question
	_ = json.Unmarshal(data, &questions)
	return questions
}
