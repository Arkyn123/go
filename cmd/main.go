package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test_service"
	"test_service/pkg/handler"
	"test_service/pkg/repository"
	"test_service/pkg/service"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt32("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			if err.Error() != "http: Server closed" {
				fmt.Printf("error occured while running http server: %s", err.Error())
			}
		}
	}()

	fmt.Println("Service is running on", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("Service shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("error on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		fmt.Printf("error on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
