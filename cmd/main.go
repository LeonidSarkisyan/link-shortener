package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
	"url-shotener-api/internal/handler"
	"url-shotener-api/internal/repositories"
	"url-shotener-api/internal/services"
	"url-shotener-api/pkg/server"
	"url-shotener-api/pkg/systems"
)

func main() {
	systems.SetupLogger()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msgf("ошибка при подключении env - файла: %s", err.Error())
	}
	config, err := systems.GetAndSetupConfig()
	if err != nil {
		log.Fatal().Msgf("ошибка при подключении конфига: %s", err.Error())
	}
	log.Info().Msg("конфиг загружен успешно")

	url := fmt.Sprintf("mongodb://%s:%s/", config.MangoDBConfig.Host, config.MangoDBConfig.Port)
	log.Info().Str("url", url).Msg("Подключение к MangoDB")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal().Msgf("ошибка при подключении к mango db: %s", err.Error())
	}
	database := client.Database(config.MangoDBConfig.Name)
	repository := repositories.NewRepositoryFromMangoDB(database)
	service := services.NewServices(repository)
	handlers := handler.NewHandler(service)
	s := new(server.Server)

	log.Info().Msgf("сервер запускается на порту: %s", viper.GetString("port"))

	go func() {
		if err = s.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatal().Msgf("серверу не удалось запуститься: %s", err.Error())
		}
	}()

	log.Info().Msgf("Сервер успешно запустился на %s порту", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	log.Info().Msg("URL-Shortener отключился")

	if err = s.Shutdown(context.Background()); err != nil {
		log.Error().Msgf("ошибка при остановки сервера: %s", err.Error())
	}

	// TODO: настроить роутер

	// TODO: запустить сервер
}
