package router

import (
	"BCScanner/config"
	"BCScanner/controllers"
	"BCScanner/domain"
	"BCScanner/infrastructure/mongo"
	"BCScanner/repository"
	"BCScanner/usecases"
	"github.com/gin-gonic/gin"
	"time"
)

func newQRGenerateRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {

	qr := repository.NewQRGenerateRepository(db, domain.CollectionQRImport)
	qc := &controllers.QRGenerateController{
		QRGenerateUseCase: usecases.NewQRGenerateUseCase(qr, timeout),
		Env:               env,
	}
	group.POST("/createQR", qc.CreateQR)
	group.GET("/fetchQR", qc.FetchQR)
	group.GET("/GetQRById", qc.GetByID)
}
