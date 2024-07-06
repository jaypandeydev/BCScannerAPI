package config

import (
	"BCScanner/infrastructure/mongo"
	//"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	config := &Application{}
	config.Env = NewEnv()
	config.Mongo = NewMongoDatabase(config.Env)
	return *config
}

func (config *Application) CloseDBConnection() {
	CloseMongoDBConnection(config.Mongo)
}
