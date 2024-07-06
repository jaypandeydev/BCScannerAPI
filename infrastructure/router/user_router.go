// infrastructure/router/router.go
package router

import (
	"BCScanner/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUserByID)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
	router.GET("/users", userController.GetAllUsers)

	return router
}
