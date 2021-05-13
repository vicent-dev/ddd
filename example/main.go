package example

import (
	ping "ddd-go/example/ping/Infrastructure/http"
	"ddd-go/http/response"
	"ddd-go/serviceProvider"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
}

//at bigger project the configuration of http should go on general
func start(sp *serviceProvider.ServiceProvider) {
	port := os.Getenv("API_PORT")

	r := mux.NewRouter()
	r.Use(response.JsonMiddleware)

	(ping.NewHandler(sp)).SetEndpoints(r)

	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatalf("can't start http server at port %s", port)
	}
	log.Printf("listening from port %s\n", port)
}
