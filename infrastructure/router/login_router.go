package router

import (
	"BCScanner/config"
	"BCScanner/controllers"
	"BCScanner/domain"
	"BCScanner/infrastructure/mongo"
	"BCScanner/repository"
	"BCScanner/usecases"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary		Logs user into the system
// @Description Responds access and refresh token as JSON post login.
// @Tags        NewLoginRouter
// @Accept  	json
// @Produce     json
// @Param   	email  query  string  true  "the user email for login"
// @Param   	password  query  string  true  "the user password for login"
// @Success     200  {string} domain
// @Router      /login [post]
func NewLoginRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controllers.LoginController{
		LoginUsecase: usecases.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
