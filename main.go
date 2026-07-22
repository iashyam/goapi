package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't laod .env")
	}

	// googeApiKey := os.Getenv("GOOGLE_API_KEY")
	model := "gemini-flash-latest"
	ctx := context.Background()
	config := &genai.GenerateContentConfig{Temperature: genai.Ptr[float32](0.5)}
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal("Can't create clinet")
	}
	chat, err := client.Chats.Create(ctx, model, config, nil)
	if err != nil {
		log.Fatal("Can't create chat")
	}
	result, err := chat.SendMessage(ctx, genai.Part{Text: "Why is the sky blue?, answer in 100 words"})
	if err != nil {
		log.Fatal("Can't ask questions")
	}
	fmt.Println(result.Text())

}
