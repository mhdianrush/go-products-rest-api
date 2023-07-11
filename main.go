package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-products-rest-api/config"
	"github.com/mhdianrush/go-products-rest-api/controllers"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/products", controllers.Index).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", controllers.Find).Methods(http.MethodGet)
	r.HandleFunc("/product", controllers.Create).Methods(http.MethodPost)
	r.HandleFunc("/product/{id}", controllers.Update).Methods(http.MethodPut)
	r.HandleFunc("/product", controllers.Delete).Methods(http.MethodDelete)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.Println("Server Running on Port 8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
