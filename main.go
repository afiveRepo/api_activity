package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatal("file .env not found")
	}

	r := gin.Default()

	r.GET("/bb", controller.FindAll)
	r.GET("/bb/:id", controller.FindDetail)
	r.POST("/bb", controller.Tambah)
	r.PUT("/bb/:id", controller.Update)
	r.DELETE("/bb/:id", controller.Hapus)

	r.Run("localhost:3030")
}
