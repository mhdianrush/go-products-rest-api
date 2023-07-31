package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-products-rest-api/config"
	"github.com/mhdianrush/go-products-rest-api/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()

	routes := mux.NewRouter()

	routes.HandleFunc("/products", controllers.Index).Methods(http.MethodGet)
	routes.HandleFunc("/product/{id}", controllers.Find).Methods(http.MethodGet)
	routes.HandleFunc("/product", controllers.Create).Methods(http.MethodPost)
	routes.HandleFunc("/product/{id}", controllers.Update).Methods(http.MethodPut)
	routes.HandleFunc("/product", controllers.Delete).Methods(http.MethodDelete)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: routes,
	}
	if err = server.ListenAndServe(); err != nil {
		logger.Printf("failed connect to server %s", err.Error())
	}

	logger.Printf("server running on port %s", os.Getenv("SERVER_PORT"))
}
