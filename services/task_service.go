package services

import (
	"errors"
	"github.com/ak-yudha/crud-gin/models"
	"github.com/ak-yudha/crud-gin/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskService interface {
	CreateTask(ctx *gin.Context, task *models.Task) error
	GetTasks() ([]models.Task, error)
	GetTaskByID(id int) (*models.Task, error)
	UpdateTask(id int, newTask *models.Task) error
	DeleteTask(id int) error
}

type TaskServiceImpl struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(taskRepository repositories.TaskRepository) TaskService {
	return &TaskServiceImpl{taskRepository}
}

func (s *TaskServiceImpl) CreateTask(ctx *gin.Context, task *models.Task) error {
	if string(task.UserId) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name, Email, and Password are required"})
		return nil
	}

	createTask, err := s.taskRepository.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return err
	}

	ctx.JSON(http.StatusCreated, createTask)
	return nil
}

func (s *TaskServiceImpl) GetTasks() ([]models.Task, error) {
	tasks, err := s.taskRepository.GetTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskServiceImpl) GetTaskByID(id int) (*models.Task, error) {
	return s.taskRepository.GetTaskByID(id)
}

func (s *TaskServiceImpl) UpdateTask(id int, newTask *models.Task) error {
	return s.taskRepository.UpdateTask(id, newTask)
}

func (s *TaskServiceImpl) DeleteTask(id int) error {
	err := s.taskRepository.DeleteTask(id)
	if err != nil {
		return errors.New("failed to delete task")
	}
	return nil
}
