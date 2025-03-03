package main

import (
	"test/internal/database"
	"test/internal/handlers"
	"test/internal/taskService"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := &taskService.TaskRepository{}
	service := &taskService.TaskService{Repository: repo}

	e := echo.New()

	api := e.Group("/api")
	{
		api.GET("/task", handlers.GetHandler(service))
		api.GET("/task/:id", handlers.GetWIDHandler(service))
		api.POST("/task", handlers.PostHandler(service))
		api.PATCH("/task/:id", handlers.PatchHandler(service))
		api.DELETE("/task/:id", handlers.DeleteHandler(service))
	}
	e.Start(":8080")
}
