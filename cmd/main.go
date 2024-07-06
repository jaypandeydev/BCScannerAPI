// cmd/main.go
package main

import (
	//"context"

	"time"

	"BCScanner/config"
	"BCScanner/infrastructure/router"

	// "BCScanner/controllers"
	// "BCScanner/infrastructure/database"
	// "BCScanner/infrastructure/router"

	"github.com/gin-gonic/gin"
	// "BCScanner/usecases"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	app := config.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	router.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)

	// //create a mongo client
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://jay:mgthyd67@cluster0.nnldyiv.mongodb.net/"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //ping the mongo server to make sure the connection is successfull
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// db := client.Database("JaytestDB")

	// // Initialize repository, usecase, and controller
	// userRepo := database.NewUserRepository(db)
	// userUsecase := usecases.NewUserUsecase(userRepo)
	// userController := controllers.NewUserController(userUsecase)

	// // Setup router
	// //r = user_router.SetupRouter(userController)
	// r := router.SetupRouter(userController)

	// // Run the server
	// r.Run()
}
