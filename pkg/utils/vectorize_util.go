package utils

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func VectorizeText(client *openai.Client, text string) ([]float32, error) {
	resp, err := client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: []string{text},
			Model: "text-embedding-3-small",
		},
	)

	if err != nil {
		fmt.Println("エラーです", err)
		return nil, err
	}

	return resp.Data[0].Embedding, nil
}
