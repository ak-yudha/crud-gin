package repositories

import (
	"database/sql"
	"github.com/ak-yudha/crud-gin/models"
)

type TaskRepository interface {
	CreateTask(task *models.Task) (int64, error)
	GetTaskByID(taskID int) (*models.Task, error)
	GetTasks() ([]models.Task, error)
	UpdateTask(id int, task *models.Task) error
	DeleteTask(taskID int) error
	GetTaskByDescription(description string) (*models.Task, error)
}

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &TaskRepositoryImpl{db}
}

// CreateTask adds a new task to the database
func (r *TaskRepositoryImpl) CreateTask(task *models.Task) (int64, error) {
	result, err := r.db.Exec("INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES (?, ?, NOW(), NOW())", task.Title, task.Description)
	if err != nil {
		return 0, err
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return taskID, nil
}

// GetTaskByID retrieves a task by their ID
func (r *TaskRepositoryImpl) GetTaskByID(taskID int) (*models.Task, error) {
	task := &models.Task{}
	err := r.db.QueryRow("SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = ?", taskID).
		Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetTasks retrieves all tasks from the database
func (r *TaskRepositoryImpl) GetTasks() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		task := models.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// UpdateTask updates task information in the database
func (r *TaskRepositoryImpl) UpdateTask(id int, task *models.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, updated_at = NOW() WHERE id = ?", task.Title, task.Description, id)
	return err
}

// DeleteTask deletes a task from the database
func (r *TaskRepositoryImpl) DeleteTask(taskID int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", taskID)
	return err
}

// GetTaskByDescription retrieves a task by their Description
func (r *TaskRepositoryImpl) GetTaskByDescription(description string) (*models.Task, error) {
	task := &models.Task{}
	err := r.db.QueryRow("SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE description = ?", description).
		Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}
