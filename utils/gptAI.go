package utils

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
)

func GptAI(ask string) string {
	config := openai.DefaultConfig("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")
	proxyUrl, err := url.Parse("http://localhost:7890")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: ask,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}
	response := resp.Choices[0].Message.Content
	//fmt.Println(response)
	return response
}
