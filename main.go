package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// -----------GET----Handler-----------------\\
func GetHandler(c echo.Context) error {
	var tasks []Task

	if err := DB.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not find the task",
		})
	}

	return c.JSON(http.StatusOK, &tasks)
}

// -----------GET-Handler--with----ID--search---\\
func GetWIDHandeler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "invalid ID",
		})
	}

	var tasks []Task

	if err := DB.Model(&Task{}).Where("id = ?", id).Find(&tasks).Error; err != nil {
		return c.JSON(echo.ErrBadRequest.Code, Response{
			Status:  "Error",
			Message: "Could not find the task",
		})
	}
	return c.JSON(http.StatusOK, &tasks)
}

// -----------POST---Handler-----------------\\
func PostHandler(c echo.Context) error {
	var tasks Task
	if err := c.Bind(&tasks); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the task",
		})
	}

	if err := DB.Create(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not create the task",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Task was succsessfully created",
	})
}

// -----------PATCH--Handler-----------------\\
func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "invalid ID",
		})
	}
	var updatedTask Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	if err := DB.Model(&Task{}).Where("id = ?", id).Update("task", updatedTask.Task).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the task",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Task was updated successfully",
	})
}

// -----------DELETE-Handler-----------------\\
func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "invalid ID",
		})
	}

	if err := DB.Delete(&Task{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not delete task",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Task was successfully deleted",
	})
}

// ------------func-----MAIN-------------------------\\
func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

	e := echo.New()

	api := e.Group("/api")
	{
		api.GET("/task", GetHandler)
		api.POST("/task", PostHandler)
		api.PATCH("/task/:id", PatchHandler)
		api.DELETE("/task/:id", DeleteHandler)
		api.GET("/task/:id", GetWIDHandeler)
	}

	e.Start(":8080")
}
