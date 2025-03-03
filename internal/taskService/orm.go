package taskService

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone string `json:"is_done"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
