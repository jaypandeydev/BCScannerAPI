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

func NewRefreshTokenRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controllers.RefreshTokenController{
		RefreshTokenUsecase: usecases.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
