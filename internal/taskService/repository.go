package taskService

import (
	"test/internal/database"
)

// репозиторий для работы с задачами
type TaskRepository struct{}

// -----Get--all--tasks-------------\\
func (r *TaskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := database.DB.Find(&tasks)
	return tasks, result.Error
}

// ------------Get--task--by--ID------------------------------\\
func (r *TaskRepository) GetTaskById(id uint) (*Task, error) {
	var task Task
	result := database.DB.First(&task, id)
	return &task, result.Error
}

// -------------Create--a---new--task----------------------------\\
func (r *TaskRepository) CreateTask(task *Task) error {
	return database.DB.Create(&task).Error
}

// ---------Update--the----task--------------------------------------\\
func (r *TaskRepository) UpdateTask(id uint, updatedTask *Task) error {
	return database.DB.Model(&Task{}).Where("id = ?", id).Updates(updatedTask).Error
}

// ------------Delete------the---task----------------------------------------\\
func (r *TaskRepository) DeleteTask(id uint) error {
	return database.DB.Delete(&Task{}, id).Error
}
