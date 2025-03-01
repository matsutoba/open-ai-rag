package main

import (
	"fmt"
	"open-ai-rag/internal/config"
	// imported as openai
)

func main() {
	config.LoadEnv()

	apiKey := config.GetEnv("OPEN_AI_API_KEY", "")
	fmt.Println(apiKey)
}
