package services

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func ConvertTextToSpeech(text string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:  "text-to-speech-001",
			Prompt: text,
		},
	)
	if err != nil {
		return "", err
	}

	audioFile := "output.mp3"
	fmt.Println(resp)
	// err = os.WriteFile(audioFile, resp.Choices[0].Text, 0644)
	// if err != nil {
	// 	return "", err
	// }

	return audioFile, nil
}
