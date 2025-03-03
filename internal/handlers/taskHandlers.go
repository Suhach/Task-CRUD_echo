package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"test/internal/taskService"
)

// --------------------GET--ALL--TASKS-----------------------------\\
func GetHandler(service *taskService.TaskService) echo.HandlerFunc {
	return func(c echo.Context) error {
		tasks, err := service.GetAllTasks()
		if err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "Could not find the task",
			})
		}
		return c.JSON(http.StatusOK, tasks)
	}
}
//--------------------GET--task---by--ID----------------------------\\
func GetWIDHandler(service *taskService.TaskService) echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "invalid ID",
			})
		}

		task, err := service.GetTaskByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "Could not find the task",
			})
		}
		return c.JSON(http.StatusOK, task)
	}
}
//-----------------------Create----------task----------------------\\
func PostHandler(service *taskService.TaskService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task taskService.Task
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "Could not add the task",
			})
		}

		if err := service.CreateTask(&task); err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "Could not create the task",
			})
		}

		return c.JSON(http.StatusOK, taskService.Response{
			Status:  "Success",
			Message: "Task was successfully created",
		})
	}
}
//-------------------------Update----task-----------------------------\\
func PatchHandler(service *taskService.TaskService) echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "invalid ID",
			})
		}

		var updatedTask taskService.Task
		if err := c.Bind(&updatedTask); err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "Invalid input",
			})
		}

		if err := service.UpdateTask(uint(id), &updatedTask); err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "Could not update the task",
			})
		}

		return c.JSON(http.StatusOK, taskService.Response{
			Status:  "Success",
			Message: "Task was updated successfully",
		})
	}
}
//-------------------------DELETE--task------------------------------\\
func DeleteHandler(service *taskService.TaskService) echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "invalid ID",
			})
		}

		if err := service.DeleteTask(uint(id)); err != nil {
			return c.JSON(http.StatusBadRequest, taskService.Response{
				Status:  "Error",
				Message: "Could not delete task",
			})
		}

		return c.JSON(http.StatusOK, taskService.Response{
			Status:  "Success",
			Message: "Task was successfully deleted",
		})
	}
}
