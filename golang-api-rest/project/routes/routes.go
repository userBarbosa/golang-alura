package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/userbarbosa/golang-alura/golang-api-rest/project/v2/controllers"
	"github.com/userbarbosa/golang-alura/golang-api-rest/project/v2/middleware"
)

var (
	ALLOWED_ORIGINS = []string{"*"}
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ResponseHeaders)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("Put")

	apiPort := os.Getenv("API_PORT")
	log.Fatal(http.ListenAndServe(":"+apiPort, handlers.CORS(handlers.AllowedOrigins(ALLOWED_ORIGINS))(r)))
}
