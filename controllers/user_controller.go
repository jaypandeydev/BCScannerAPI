// controllers/user_controller.go
package controllers

import (
	//"errors"
	"net/http"

	"BCScanner/config"
	"BCScanner/domain"

	//"BCScanner/usecases"

	//"BCScanner/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *config.Env
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

func (ctrl *UserController) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.UserUsecase.Create(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.UserUsecase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) Fetch(c *gin.Context) {
	var user []domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctrl.UserUsecase.Fetch(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetByEmail(c *gin.Context) {
	var request domain.User
	users, err := ctrl.UserUsecase.GetByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// func (ctrl *UserController) DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := ctrl.UserUsecase.DeleteUser(id); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }
