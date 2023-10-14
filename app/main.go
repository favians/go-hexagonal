package main

import (
	"context"
	"fmt"
	"go-hexagonal/config"
	"go-hexagonal/modules/mongodb"

	"go-hexagonal/api"
	messagesController "go-hexagonal/api/v1/messages"
	userController "go-hexagonal/api/v1/user"
	messagesService "go-hexagonal/business/messages"
	userService "go-hexagonal/business/user"
	messagesRepository "go-hexagonal/modules/messages"
	userRepository "go-hexagonal/modules/user"

	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
)

func newDatabaseConnection(config *config.AppConfig) *mongo.Database {
	uri := "mongodb://"

	if config.AppEnvironment == "prod" {
		uri = "mongodb+srv://"
	}

	if config.DbUsername != "" {
		uri = fmt.Sprintf("%s%v:%v@", uri, config.DbUsername, config.DbPassword)
	}

	if config.AppEnvironment == "prod" {
		uri = fmt.Sprintf("%s%v/factura?retryWrites=true&w=majority",
			uri,
			config.DbAddress,
		)
	} else {
		uri = fmt.Sprintf("%s%v:%v/?connect=direct",
			uri,
			config.DbAddress,
			config.DbPort,
		)
	}

	db, err := mongodb.NewDatabaseConnection(uri, config.DbName)

	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	dbConnection := newDatabaseConnection(config)

	//create echo http
	e := echo.New()

	//initiate messages
	messagesRepo := messagesRepository.NewMongoDBRepository(dbConnection)
	messagesService := messagesService.NewService(messagesRepo)
	messagesController := messagesController.NewController(messagesService)

	//initiate users
	userRepo := userRepository.NewMongoDBRepository(dbConnection)
	userService := userService.NewService(userRepo)
	userController := userController.NewController(userService)

	//register paths
	api.RegisterPaths(e, messagesController, userController)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%d", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
