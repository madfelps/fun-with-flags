package routes

import (
	"fun-w-flags/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.TodosPaises)
	log.Fatal(http.ListenAndServe(":5200", r))
	//r.RUN(":5200", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}
