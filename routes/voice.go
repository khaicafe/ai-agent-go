package routes

import (
	"ai-agent-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VoiceRoutes(router *gin.Engine) {
	router.POST("/voice", func(c *gin.Context) {
		file, err := c.FormFile("audio")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File không hợp lệ"})
			return
		}

		// Lưu file tạm thời
		filePath := "./temp/" + file.Filename
		c.SaveUploadedFile(file, filePath)

		// Chuyển giọng nói thành văn bản
		text, err := services.ConvertSpeechToText(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi STT"})
			return
		}

		// Gửi đến AI xử lý
		aiResponse, err := services.ProcessAIRequest(text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi AI"})
			return
		}

		// Tạo giọng nói từ phản hồi AI
		audioFile, err := services.ConvertTextToSpeech(aiResponse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi TTS"})
			return
		}

		c.File(audioFile)
	})
}
