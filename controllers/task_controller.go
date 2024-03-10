package controllers

import (
	"fmt"
	"github.com/ak-yudha/crud-gin/models"
	"github.com/ak-yudha/crud-gin/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TaskController interface {
	CreateTask(ctx *gin.Context)
	GetTasks(ctx *gin.Context)
	GetTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

type TaskControllerImpl struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) TaskController {
	return &TaskControllerImpl{taskService}
}

func (c *TaskControllerImpl) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.taskService.CreateTask(ctx, &task); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to register task"})
		return
	}

	ctx.JSON(201, gin.H{"id": task.ID})
}

func (c *TaskControllerImpl) GetTasks(ctx *gin.Context) {
	tasks, err := c.taskService.GetTasks()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, tasks)
}

func (c *TaskControllerImpl) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	task, err := c.taskService.GetTaskByID(taskId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(200, task)
}

func (c *TaskControllerImpl) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var task *models.Task
	if err := ctx.BindJSON(&task); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	err = c.taskService.UpdateTask(taskId, task)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(200, err)
}

func (c *TaskControllerImpl) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = c.taskService.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Task deleted successfully"})
}
