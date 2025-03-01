package samples

import (
	"fmt"
	"open-ai-rag/pkg/utils"

	"github.com/sashabaranov/go-openai"
)

func WebTextQandA(client *openai.Client, url string, question string) {
	text := utils.GetText(url)
	chunkedTexts := utils.SplitIntoChunks(text, 300, 50)

	// 回答のベクトル化
	var vectors [][]float32
	for _, text := range chunkedTexts {
		vector, err := utils.VectorizeText(client, text)
		if err != nil {
			fmt.Println("回答のベクトル化エラーです", err)
		}
		vectors = append(vectors, vector)
	}

	// 質問のベクトル化
	question_vector, err := utils.VectorizeText(client, question)
	if err != nil {
		fmt.Println("質問のベクトル化エラーです", err)
	}

	similarDocument := utils.FindMostSimilarDocument(question_vector, vectors, chunkedTexts)

	answer := utils.AskQuestion(client, question, similarDocument)

	fmt.Printf("\n")
	fmt.Printf("[回答]\n")
	fmt.Println(answer)
}
