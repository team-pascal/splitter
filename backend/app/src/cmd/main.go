package main

import (
	"fmt"
	"os"
	"splitter/internal/routers"
	"splitter/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("This is main.go")
	db := database.InitDB()

	r := gin.Default()
	routers.SetupRouter(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)

}
