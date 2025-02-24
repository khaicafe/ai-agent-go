package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Hàm 1: Lấy thời gian hiện tại
func GetCurrentTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf(`{"time": "%s"}`, currentTime)
}

// Hàm 2: Tính toán biểu thức đơn giản (vd: "3 + 5")
func CalculateExpression(expression string) string {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return `{"error": "Invalid format. Example: '3 + 5'" }`
	}

	num1, err1 := strconv.ParseFloat(parts[0], 64)
	operator := parts[1]
	num2, err2 := strconv.ParseFloat(parts[2], 64)

	if err1 != nil || err2 != nil {
		return `{"error": "Invalid numbers in expression."}`
	}

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return `{"error": "Division by zero"}`
		}
		result = num1 / num2
	default:
		return `{"error": "Invalid operator. Use +, -, *, /"}`
	}

	return fmt.Sprintf(`{"result": %.2f}`, result)
}

// Hàm 3: Lấy thời tiết
func GetWeather(location string) string {
	return fmt.Sprintf(`{"weather": "Sunny, 25°C in %s"}`, location)
}

// Xử lý lời gọi hàm từ AI
func HandleFunctionCall(functionName string, args json.RawMessage) (string, error) {
	switch functionName {
	case "get_time":
		return GetCurrentTime(), nil
	case "calculate":
		var params struct {
			Expression string `json:"expression"`
		}
		json.Unmarshal(args, &params)
		return CalculateExpression(params.Expression), nil
	case "get_weather":
		var params struct {
			Location string `json:"location"`
		}
		json.Unmarshal(args, &params)
		return GetWeather(params.Location), nil
	default:
		return `{"error": "Unknown function"}`, nil
	}
}
