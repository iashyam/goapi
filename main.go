package main

import (
	"context"
	"fmt"

	"github.com/iashyam/goapi/api"
	
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func askAI(question string) (string, error) {
	model := "gemini-flash-latest"
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	config := &genai.GenerateContentConfig{Temperature: genai.Ptr[float32](0.5)}
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return "", err
	}
	chat, err := client.Chats.Create(ctx, model, config, nil)
	if err != nil {
		return "", err
	}
	result, err := chat.SendMessage(ctx, genai.Part{Text: question})
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}

func main() {
	// question := "Why doesn't the sun run out of fuel? answer in 100 words."
	// answer, err := askAI(question)
	// if err != nil {
	// 	fmt.Printf("Error while asking question %v", err)
	// }
	answer := "hello world"
	fmt.Println(answer)
	PrintHelloWorld()
}
