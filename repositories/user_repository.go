package repositories

import (
	"database/sql"
	"github.com/ak-yudha/crud-gin/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (int64, error)
	GetUserByID(userID int) (*models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id int, user *models.User) error
	DeleteUser(userID int) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

// CreateUser adds a new user to the database
func (r *UserRepositoryImpl) CreateUser(user *models.User) (int64, error) {
	result, err := r.db.Exec("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, NOW(), NOW())", user.Name, user.Email)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepositoryImpl) GetUserByID(userID int) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?", userID).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers retrieves all users from the database
func (r *UserRepositoryImpl) GetUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, password, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser updates user information in the database
func (r *UserRepositoryImpl) UpdateUser(id int, user *models.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ?, email = ?, password = ?, updated_at = NOW() WHERE id = ?", user.Name, user.Email, id)
	return err
}

// DeleteUser deletes a user from the database
func (r *UserRepositoryImpl) DeleteUser(userID int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", userID)
	return err
}

// GetUserByEmail retrieves a user by their Email
func (r *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
