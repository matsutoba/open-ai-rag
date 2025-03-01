package main

import (
	"open-ai-rag/pkg/utils"
	"open-ai-rag/samples"

	"github.com/sashabaranov/go-openai"
	// imported as openai
)

func main() {
	utils.LoadEnv()
	apiKey := utils.GetEnv("OPEN_AI_API_KEY", "")
	client := openai.NewClient(apiKey)

	url := "質問をしたいページのURL"
	question := "質問文"
	samples.WebTextQandA(client, url, question)
}
