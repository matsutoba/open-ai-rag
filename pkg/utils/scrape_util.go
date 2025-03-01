package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetText(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("ページ取得エラーです", err)
		return ""
	}
	defer response.Body.Close()

	// HTMLをgoqueryに渡して解析
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("Bodyの読み取りエラーです", err)
		return ""
	}

	// <script>、 <style>, <head>タグを除外
	doc.Find("script, style, head").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// HTMLからテキスト部分を取得
	text := doc.Text()

	// 改行やタブを削除し、余分な空白を1つにまとめる
	text = strings.ReplaceAll(text, "\n", " ")     // 改行をスペースに置き換え
	text = strings.ReplaceAll(text, "\t", " ")     // タブをスペースに置き換え
	text = strings.Join(strings.Fields(text), " ") // 余分な空白を1つにまとめる

	return text
}

/*
大きなテキストをLLMに送らないようにするため、指定サイズのチャンクに分割
*/
func SplitIntoChunks(text string, chunkSize int, overlapSize int) []string {
	var chunks []string
	runes := []rune(text) // 文字単位でスライス化
	textLength := len(runes)

	// チャンク分割処理（文字単位）
	for i := 0; i < textLength; i += chunkSize - overlapSize {
		// チャンクの終端位置を決める（文字単位）
		end := i + chunkSize
		if end > textLength {
			end = textLength
		}

		// チャンクを追加
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}
