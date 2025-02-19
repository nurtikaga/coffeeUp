package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	coffeeup "github.com/nurtikaga/coffeeUp"
	"github.com/nurtikaga/coffeeUp/pkg/handler"
	"github.com/nurtikaga/coffeeUp/pkg/repository"
	"github.com/nurtikaga/coffeeUp/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Init Configs wrong: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Password from env wrong: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   viper.GetString("db.dbname"),
		Sslmode:  viper.GetString("db.sslmode"),
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	if err != nil {
		log.Fatalf("Couldnt connect to DB ", err.Error())
	}

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
