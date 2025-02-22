package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	coffeeup "github.com/nurtikaga/coffeeUp"
	"github.com/nurtikaga/coffeeUp/pkg/handler"
	"github.com/nurtikaga/coffeeUp/pkg/repository"
	"github.com/nurtikaga/coffeeUp/pkg/service"
	"github.com/rs/cors" // Import the CORS package
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfigs(); err != nil {
		logrus.Fatalf("Init Configs wrong: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Password from env wrong: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   viper.GetString("db.dbname"),
		Sslmode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Couldnt connect to DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	// Initialize CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your React frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true, // Enable debug logging for CORS
	})

	// Wrap your router with the CORS middleware
	router := handler.InitRoutes()
	handlerWithCORS := corsMiddleware.Handler(router)

	srv := new(coffeeup.Server)
	if err := srv.Run(viper.GetString("port"), handlerWithCORS); err != nil {
		logrus.Fatalf("Server doesnt run: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
