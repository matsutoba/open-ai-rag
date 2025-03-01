package utils

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

func AskQuestion(client *openai.Client, question string, informations []string) string {
	informationContexts := strings.Join(informations, "")

	prompt :=
		"以下の質問に以下の情報をベースにして答えてください。\n" +
			"[質問" + strconv.FormatInt(time.Now().UnixNano(), 10) + "]\n" +
			question + "\n" +
			"\n" +
			"[情報]\n" +
			informationContexts

	completionRequest := openai.CompletionRequest{
		Model:     openai.GPT3Dot5TurboInstruct, // CreateCompletionで使えるモデル
		Prompt:    prompt,
		MaxTokens: 200,
	}

	completionResponse, err := client.CreateCompletion(
		context.Background(),
		completionRequest,
	)
	if err != nil {
		fmt.Println("Completion生成のエラー", err)
		return ""
	}

	fmt.Printf("[プロンプト]\n%v\n\n", prompt)

	return completionResponse.Choices[0].Text
}
