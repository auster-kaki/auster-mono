package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	/*
		月毎に５ドル分のクレジットがあります。
		それ以上は、従量課金でした。
		目安の最大で５ドルで1000リクストが可能みたいです。
		が利用するモデルや、リクエスト時のトークンの数によって、リクエスト可能回数が変わります。

		５ドルを超えるリクエストを投げると、エラーが返ってくる想定です。
		ですが、apiで現状のクレジットの確認ができないです。管理画面を確認する必要があります。

		1分単位のリクエスト制限がありますが、厳しくなかったので、気にしなくて良さそうです。
		50リクエスト/分くらいででした。
	*/

	apiKey := "YOUR_API_KEY_HERE"
	model := "llama-3.1-sonar-large-128k-online"
	messages := []Message{
		{
			Role:    "system",
			Content: "あなたは日記作成を支援するAIアシスタントです。与えられた情報をもとに、指定された文字数でタイトルと本文を生成してください。",
		},
		{
			Role: "user",
			Content: "以下の設定で日記のタイトルと本文を考えてください：\n" +
				"- タイトル：5〜20文字程度\n" +
				"- 本文：100〜200文字程度\n" +
				"- 趣味：釣り\n" +
				"- 体験：釣り体験\n" +
				"- 内容：早朝に漁船に乗り釣りを行った\n\n" +
				"タイトルと本文を分けて出力してください。",
		},
	}

	res, err := perplexityRequest(apiKey, model, messages)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(res)

	/* 実際のレスポンスです
	&main.PerplexityResponse{
	  ID:      "19b77111-282b-4c71-ae25-86cb4ea06d75",
	  Model:   "llama-3.1-sonar-large-128k-online",
	  Created: 1731655879,
	  Usage:   main.UsageInfo{
	    PromptTokens:     134,
	    CompletionTokens: 201,
	    TotalTokens:      335,
	  },
	  Citations: []string{
	    "https://kohama-tour.com/plan/fishing607",
	    "https://www.jalan.net/kankou/spt_guide000000154028/activity/l00005B6D5/?showplan=ichiran_guide&asobiKbn=1",
	    "https://sinseimaru.com/column/detail/20231230/",
	    "https://www.okinawauminchu.com/cruise.html",
	    "https://www.asoview.com/leisure/act0172/",
	  },
	  Object:  "chat.completion",
	  Choices: []main.ChoiceInfo{
	    main.ChoiceInfo{
	      Index:        0,
	      FinishReason: "stop",
	      Message:      main.MessageInfo{
	        Role:    "assistant",
	        Content: "## タイトル\n早朝の漁船釣り体験\n\n## 本文\n今日は早朝、小浜島での漁船釣り体験をしました。5:30に小浜港で集合し、日の出を目指して出航。早朝は魚が活発に動く時間帯で、釣果が期待できました。ガイドの指導を受けながら、五目釣りを楽しむことができました。釣り初心者でも手軽に楽しめるように、釣り竿や餌、ライフジャケットまで全て用意されていました。日の出を見ながら魚を釣る楽しさは言葉で説明しづらいほどでした。釣った魚は持ち帰りできたので、夕食に楽しむ予定です。早朝の海での冒険は本当に素晴らしい体験でした。",
	      },
	      Delta: main.DeltaInfo{
	        Role:    "assistant",
	        Content: "",
	      },
	    },
	  },
	}
	*/
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func perplexityRequest(apiKey, model string, messages []Message) (*PerplexityResponse, error) {
	url := "https://api.perplexity.ai/chat/completions"

	requestBody := Request{
		Model:    model,
		Messages: messages,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res PerplexityResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// PerplexityResponse はAPIレスポンス全体を表す構造体
type PerplexityResponse struct {
	ID        string       `json:"id"`
	Model     string       `json:"model"`
	Created   int64        `json:"created"`
	Usage     UsageInfo    `json:"usage"`
	Citations []string     `json:"citations"`
	Object    string       `json:"object"`
	Choices   []ChoiceInfo `json:"choices"`
}

// CreatedTime はCreatedフィールドをtime.Time型として返すメソッド
func (pr *PerplexityResponse) CreatedTime() time.Time {
	return time.Unix(pr.Created, 0)
}

// UsageInfo はトークン使用量の情報を表す構造体
type UsageInfo struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChoiceInfo は生成された選択肢の情報を表す構造体
type ChoiceInfo struct {
	Index        int         `json:"index"`
	FinishReason string      `json:"finish_reason"`
	Message      MessageInfo `json:"message"`
	Delta        DeltaInfo   `json:"delta"`
}

// MessageInfo は生成されたメッセージの情報を表す構造体
type MessageInfo struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// DeltaInfo はストリーミング応答時のデルタ情報を表す構造体
type DeltaInfo struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
