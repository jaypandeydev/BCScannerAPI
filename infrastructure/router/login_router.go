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

func NewLoginRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controllers.LoginController{
		LoginUsecase: usecases.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
