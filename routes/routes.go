package routes

import (
	"client/controllers"
	"client/middleware"
	"client/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/_healthcheck", controllers.Home).Methods("GET")
	r.HandleFunc("/create-user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+services.GetEnv("CLIENT_SERVER_PORT"),
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r),
	))
}
