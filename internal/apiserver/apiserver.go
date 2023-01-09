package apiserver

import (
    "os"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
    "geocoder/internal/handlers"
)

type APIServer struct {
    port string
    logger *logrus.Logger
    router *mux.Router
}

func APIServerFactory() *APIServer {
    port, exists := os.LookupEnv("PORT")

    if !exists {
	     port = "80"
    }

    router := mux.NewRouter()
    router.Use(AuthenticationMiddleware)

    return &APIServer {
        port: port,
        logger: logrus.New(),
        router: router,
    }
}

func (server *APIServer) Run() error {
    if err := server.configureLogger(); err != nil {
        return err
    }

    server.configureRouter()

    server.logger.Info("Starting api server")

    return http.ListenAndServe(":" + server.port, server.router)
}

func (server *APIServer) configureLogger() error {
    level, err := logrus.ParseLevel("debug")

    if err != nil {
        return err
    }

    server.logger.SetLevel(level)

    return nil
}

func (server *APIServer) configureRouter() {
    server.router.HandleFunc("/geocode", handlers.GeocodeHandler)
    server.router.HandleFunc("/reverse-geocode", handlers.ReverseGeocodeHandler)
}
