ai-agent-go/
│── main.go
│── routes/
│ ├── agent.go # API xử lý yêu cầu
│── services/
│ ├── ai_service.go # Gửi request đến ChatGPT
│ ├── functions.go # Định nghĩa các hàm có thể gọi
│── utils/
│ ├── config.go # Load API key từ .env
│── .env # API key OpenAI
