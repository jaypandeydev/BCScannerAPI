// cmd/main.go
package main

import (
	"BCScanner/config"
	_ "BCScanner/docs"
	"BCScanner/infrastructure/router"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     QRcode Scanning API
// @version         1.0
// @description     A QR Code Scanning service API in Go using Gin framework.

// @contact.name   Jay Pandey
// @contact.url    https://twitter.com/jaypandeyspeaks
// @contact.email  pandeyjp1@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

func main() {

	app := config.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	router.Setup(env, timeout, db, gin)

	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.Run(env.ServerAddress)
}
