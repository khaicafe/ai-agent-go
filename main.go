package main

import (
	"ai-agent-go/routes"
	"ai-agent-go/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	utils.LoadEnv()
	router := gin.Default()

	routes.AgentRoutes(router)

	router.Run(":8080")
}

func mainBK() {
	client := resty.New()

	apiKey := "sk-b550438cf64440deacc2c3fc06ecc0a1"
	url := "https://api.deepseek.com/v1/chat/completions" // DeepSeek endpoint

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(`{
			"model": "deepseek-chat", 
			"messages": [{"role": "user", "content": "Hello, how are you?"}]
		}`).
		Post(url)

	if err != nil {
		fmt.Println("Lỗi khi gọi API:", err)
		return
	}

	fmt.Println("Kết quả:", resp.String())
}
