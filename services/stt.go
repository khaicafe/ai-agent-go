package services

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// Speech-to-Text (STT) với OpenAI Whisper
func ConvertSpeechToText(audioFilePath string) (string, error) {
	apiURL := "https://api.openai.com/v1/audio/transcriptions"
	apiKey := os.Getenv("OPENAI_API_KEY")

	// Mở file audio
	file, err := os.Open(audioFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Chuẩn bị multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", audioFilePath)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	writer.WriteField("model", "whisper-1")
	writer.Close()

	// Gửi request đến OpenAI
	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Đọc kết quả
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
