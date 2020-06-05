package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

    "github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	router "github.com/go-chi/chi"

	"generator/server/handlers"
	"generator/server/middleware"
)

const keyENV = "APP_ENV"

var version = "dev"

func readConfig(env string) {
	if len(env) > 0 {
		env = fmt.Sprintf(".%s", env)
	}

	viper.SetConfigFile(fmt.Sprintf("./config/app%s.yml", env))
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	app := "generator"
	env := os.Getenv(keyENV)

	readConfig(env)

    logger := log.New()

	logger.Info(fmt.Sprintf("Starting %s on %s env..", app, env))

	host, port := viper.GetString("host"), viper.GetString("port")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT)

	r := router.NewRouter()

	lm := middleware.NewLog(logger, true)
	r.Use(lm.Handler)

	ping := handlers.NewPing(logger, version)
    r.Get("/ping", ping.Handler)

    hello := handlers.NewHello(logger)
	r.Get("/", hello.Handler)

    urand := handlers.NewUrandom(logger)
    r.Get("/urandom", urand.Handler)

	fbsdRand := handlers.NewFreebsdrandom(logger)
	r.Get("/fbsdrandom", fbsdRand.Handler)

	httpErr := make(chan error, 1)
	go func() {
		logger.Info(fmt.Sprintf("Started server on %s:%s..", host, port))
		httpErr <- http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r)
	}()

	select {
    case err := <-httpErr:
        logger.Error(err.Error())
    case <-stop:
        logger.Info("Stopped via signal")
    }

    logger.Info(fmt.Sprintf("Stopping %s..", app))
}
