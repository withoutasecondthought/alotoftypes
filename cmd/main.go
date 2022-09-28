package main

import (
	"alotoftypes"
	configs "alotoftypes/config"
	"alotoftypes/pkg/handler"
	"alotoftypes/pkg/repository"
	"alotoftypes/pkg/service"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	err := configs.InitConfig()
	if err != nil {
		logrus.Fatalf("Error Config initialization %s", err)
	}

	db, err := initDB()
	if err != nil {
		logrus.Fatalf("Error Database initialization %s", err)
	}

	repos := repository.NewRepository(db, viper.GetString("db.collection"))
	ser := service.NewService(repos)
	handlers := handler.NewHandler(ser)

	srv := new(alotoftypes.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err)
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initDB() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("db.uri")))
	if err != nil {
		return nil, err
	}

	ctx, cansel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cansel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(viper.GetString("db.database")), nil
}
