package taskService

type TaskService struct {
	Repository *TaskRepository
}

// ------------------GET---THE--ALL--TASKS------------------\\
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.Repository.GetAllTasks()
}

// ---------GET--THE---TASK--BY--ID---------------------------------\\
func (s *TaskService) GetTaskByID(id uint) (*Task, error) {
	return s.Repository.GetTaskById(id)
}

// ----------------CREATE--THE-TASK------------------------------\\
func (s *TaskService) CreateTask(task *Task) error {
	return s.Repository.CreateTask(task)
}

// -----------------------UPDATE--THE---TASK-------------------------\\
func (s *TaskService) UpdateTask(id uint, updatedTask *Task) error {
	return s.Repository.UpdateTask(id, updatedTask)
}

// --------------DELETE--THE-TASK------------------------------\\
func (s *TaskService) DeleteTask(id uint) error {
	return s.Repository.DeleteTask(id)
}
