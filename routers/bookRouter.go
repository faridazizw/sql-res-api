package router

import (
	"github.com/gorilla/mux"
	"sql-rest-api/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/books", controllers.AmbilSemuaBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controllers.AmbilBuku).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books", controllers.TmbhBuku).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBuku).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controllers.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}
