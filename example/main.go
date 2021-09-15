package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vicent-dev/ddd/http/response"
	"github.com/vicent-dev/ddd/serviceProvider"

	"example/ping/Infrastructure"

	ping "example/ping/Infrastructure/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	start(Infrastructure.NewServiceProvider())
}

func start(sp *serviceProvider.ServiceProvider) {
	port := os.Getenv("API_PORT")

	r := mux.NewRouter()
	r.Use(response.JsonMiddleware)

	(ping.NewHandler(sp)).SetEndpoints(r)

	log.Printf("Running server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
