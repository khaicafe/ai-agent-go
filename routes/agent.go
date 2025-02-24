package routes

import (
	"ai-agent-go/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AgentRoutes(router *gin.Engine) {
	router.POST("/ask", func(c *gin.Context) {
		var request struct {
			Message string `json:"message"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Gọi AI Agent để xử lý
		response, err := services.ProcessAIRequest(request.Message)
		if err != nil {
			fmt.Println("🔴 AI Processing Error:", err) // Ghi log lỗi
			c.JSON(http.StatusInternalServerError, gin.H{"error": "AI processing failed", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": response})
	})
}
