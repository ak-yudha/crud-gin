package controllers

import (
	"fmt"
	"github.com/ak-yudha/crud-gin/models"
	"github.com/ak-yudha/crud-gin/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController interface {
	RegisterUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{userService}
}

func (c *UserControllerImpl) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.RegisterUser(ctx, &user); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(201, gin.H{"id": user.ID})
}

func (c *UserControllerImpl) GetUsers(ctx *gin.Context) {
	users, err := c.userService.GetUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

func (c *UserControllerImpl) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	user, err := c.userService.GetUserByID(userId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(200, user)
}

func (c *UserControllerImpl) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var user *models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	err = c.userService.UpdateUser(userId, user)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(200, err)
}

func (c *UserControllerImpl) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = c.userService.DeleteUser(userId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
