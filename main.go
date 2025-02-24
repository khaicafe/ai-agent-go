package main

import (
	"ai-agent-go/routes"
	"ai-agent-go/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	router := gin.Default()

	routes.AgentRoutes(router)

	router.Run(":8080")
}
