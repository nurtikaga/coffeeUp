package main

import (
	"log"

	coffeeup "github.com/nurtikaga/coffeeUp"
	"github.com/nurtikaga/coffeeUp/pkg/handler"
	"github.com/nurtikaga/coffeeUp/pkg/repository"
	"github.com/nurtikaga/coffeeUp/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Something wrong: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(coffeeup.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Server doesnt run", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
