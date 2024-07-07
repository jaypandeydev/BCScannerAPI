// infrastructure/router/router.go
package router

import (
	"time"

	"BCScanner/config"
	"BCScanner/domain"
	"BCScanner/infrastructure/mongo"
	"BCScanner/repository"

	"BCScanner/controllers"
	"BCScanner/usecases"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)

	uc := &controllers.UserController{
		UserUsecase: usecases.NewUserUsecase(ur, timeout),
		Env:         env,
	}
	group.POST("/create", uc.Create)
	group.GET("/fetch", uc.Fetch)
	group.GET("/getbyemail", uc.GetByEmail)
	group.GET("/getbyid", uc.GetByID)

	//router := gin.Default()

	// router.POST("/users", userController.CreateUser)
	// router.GET("/users/:id", userController.GetUserByID)
	// router.PUT("/users/:id", userController.UpdateUser)
	// router.DELETE("/users/:id", userController.DeleteUser)
	// router.GET("/users", userController.GetAllUsers)

	// return router
}
