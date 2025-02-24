package services

import (
	"context"
	"encoding/json"
	"os"

	"github.com/sashabaranov/go-openai"
)

func ProcessAIRequest(userInput string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			// Model: openai.GPT4Turbo,
			// Model: "gpt-4o", // 🔹 Sử dụng model có quyền truy cập
			// Model: "gpt-3.5-turbo",

			Model: "gpt-4o-mini",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Bạn là một trợ lý AI có thể gọi hàm.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userInput,
				},
			},
			Tools: []openai.Tool{
				{
					Type: openai.ToolTypeFunction,
					Function: &openai.FunctionDefinition{
						Name:        "get_time",
						Description: "Get the current time",
					},
				},
				{
					Type: openai.ToolTypeFunction,
					Function: &openai.FunctionDefinition{
						Name:        "calculate",
						Description: "Calculate a mathematical expression",
						Parameters: map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"expression": map[string]interface{}{
									"type":        "string",
									"description": "Math expression, e.g., '3 + 5'",
								},
							},
							"required": []string{"expression"},
						},
					},
				},
				{
					Type: openai.ToolTypeFunction,
					Function: &openai.FunctionDefinition{
						Name:        "get_weather",
						Description: "Get the weather for a location",
						Parameters: map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"location": map[string]interface{}{
									"type":        "string",
									"description": "City name",
								},
							},
							"required": []string{"location"},
						},
					},
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	// Nếu AI gọi một Tool (Function Calling)
	if len(resp.Choices) > 0 && resp.Choices[0].Message.ToolCalls != nil {
		toolCall := resp.Choices[0].Message.ToolCalls[0]
		functionName := toolCall.Function.Name
		argsString := toolCall.Function.Arguments

		// 🔹 Chuyển đổi `argsString` thành `json.RawMessage`
		argsJSON := json.RawMessage(argsString)

		// Gọi hàm tương ứng
		result, err := HandleFunctionCall(functionName, argsJSON)
		if err != nil {
			return "", err
		}

		return result, nil
	}

	// Nếu AI không gọi hàm, trả về phản hồi bình thường
	return resp.Choices[0].Message.Content, nil
}
