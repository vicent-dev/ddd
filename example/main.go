package main

import (
	"log"
	"net/http"
	"os"

	"example/ping/Infrastructure"

	ping "example/ping/Infrastructure/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/vicent-dev/ddd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	start(Infrastructure.NewServiceProvider())
}

func start(sp *ddd.ServiceProvider) {
	port := os.Getenv("API_PORT")

	r := mux.NewRouter()
	r.Use(ddd.JsonMiddleware)

	(ping.NewHandler(sp)).SetEndpoints(r)

	log.Printf("Running server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
