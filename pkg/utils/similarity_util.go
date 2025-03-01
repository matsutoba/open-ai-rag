package utils

import (
	"fmt"
	"math"
	"sort"
)

// コサイン類似度を計算する関数
func CosineSimilarity(vec1, vec2 []float32) float32 {
	var dotProduct, normVec1, normVec2 float32

	// 内積と各ベクトルのノルムを計算
	for i := 0; i < len(vec1); i++ {
		dotProduct += vec1[i] * vec2[i]
		normVec1 += vec1[i] * vec1[i]
		normVec2 += vec2[i] * vec2[i]
	}

	// コサイン類似度の計算
	return dotProduct / (float32(math.Sqrt(float64(normVec1))) * float32(math.Sqrt(float64(normVec2))))
}

func FindMostSimilarDocument(vec1 []float32, vec2 [][]float32, documents []string) []string {
	type Data struct {
		Index      int
		Similarity float32
	}
	var similarities []Data

	// 類似度が高いドキュメントをいくつ採用するか
	const topN = 2

	for index, vector := range vec2 {
		similarity := CosineSimilarity(vec1, vector)
		fmt.Printf("%v %v\n", similarity, documents[index])
		data := Data{Index: index, Similarity: similarity}
		similarities = append(similarities, data)
	}

	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].Similarity > similarities[j].Similarity
	})

	var topDocuments []string
	for i := 0; i < topN; i++ {
		index := similarities[i].Index
		topDocuments = append(topDocuments, documents[index])
	}

	return topDocuments
}
