package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	fmt.Println(">> Alaska Small Business Development Center | go version go1.22.2 linux/amd64")

	client := openai.NewClient(
		option.WithAPIKey("My API Key"),
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("What are the main jurisdictional funding sources in Alaska?"),
		},
		Model: openai.ChatModelChatgpt4oLatest,
	})
	if err != nil {
		panic(err.Error())
	}
	
	println(chatCompletion.Choices[0].Message.Content)
}