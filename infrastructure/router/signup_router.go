package router

import (
	"time"

	"BCScanner/config"
	"BCScanner/controllers"
	"BCScanner/domain"
	"BCScanner/infrastructure/mongo"
	"BCScanner/repository"
	"BCScanner/usecases"

	"github.com/gin-gonic/gin"
)

// @Summary		Complete the signup
// @Description Return access and refresh token as JSON.
// @Tags        NewSignupRouter
// @Accept  	json
// @Produce     json
// @Param   	email  query  string  true  "the user email for login"
// @Param   	password  query  string  true  "the user password for login"
// @Param   	name  query  string  true  "the user name for login"
// @Success     200  {string}  domain.SignupResponse
// @Router      /signup [post]
func NewSignupRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controllers.SignupController{
		SignupUsecase: usecases.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
