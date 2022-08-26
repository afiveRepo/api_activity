package main

import (
	"log"
	"time"

	"api_activity/config"
	"api_activity/controller"
	"api_activity/repository"
	"api_activity/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatal("file .env not found")
	}
	db := config.ConnectionDB()
	acgRepo := repository.New(db)
	acgService := services.New(acgRepo)
	acgController := controller.New(acgService)

	todoRepo := repository.NewTodo(db)
	todoService := services.NewTodo(acgRepo, todoRepo)
	todoController := controller.NewTodo(todoService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowWildcard:    true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "GET", "HEAD", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "*", "Authorization", "Content-Disposition"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/activity-groups", acgController.FindAll)
	r.GET("/activity-groups/:id", acgController.FindByID)
	r.POST("/activity-groups", acgController.Create)
	r.PATCH("/activity-groups/:id", acgController.UpdateByID)
	r.DELETE("/activity-groups/:id", acgController.DeleteByID)

	r.GET("/todo-items", todoController.FindAll)
	r.GET("/todo-items/:id", todoController.FindByID)
	//r.GET("/todo-items",todoController.FindByGroupID)
	r.POST("/todo-items", todoController.Create)
	r.PATCH("/todo-items/:id", todoController.UpdateByID)
	r.DELETE("/todo-items/:id", todoController.DeleteByID)

	r.Run("localhost:3030")
}
