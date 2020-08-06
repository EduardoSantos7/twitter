package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/EduardoSantos7/twitter/middlew"
	"github.com/EduardoSantos7/twitter/routers"

	// "github.com/gorila/mux"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Port set and start to listen and handlers*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.DBCheck(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
