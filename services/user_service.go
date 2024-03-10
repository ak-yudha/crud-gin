package services

import (
	"errors"
	"github.com/ak-yudha/crud-gin/models"
	"github.com/ak-yudha/crud-gin/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserService interface {
	RegisterUser(ctx *gin.Context, user *models.User) error
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, newUser *models.User) error
	DeleteUser(id int) error
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (s *UserServiceImpl) RegisterUser(ctx *gin.Context, user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name, Email, and Password are required"})
		return nil
	}

	existingUser, err := s.userRepository.GetUserByEmail(user.Email)
	if err == nil || existingUser != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email is already in use"})
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return err
	}
	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()

	createUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return err
	}

	ctx.JSON(http.StatusCreated, createUser)
	return nil
}

func (s *UserServiceImpl) GetUsers() ([]models.User, error) {
	users, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) GetUserByID(id int) (*models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *UserServiceImpl) UpdateUser(id int, newUser *models.User) error {
	return s.userRepository.UpdateUser(id, newUser)
}

func (s *UserServiceImpl) DeleteUser(id int) error {
	err := s.userRepository.DeleteUser(id)
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
