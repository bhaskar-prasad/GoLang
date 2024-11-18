package router

import (
	"github.com/bhaskar-prasad/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Home).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}/watched", controller.MarkAsWatched).Methods("PUT")	
	router.HandleFunc("/api/delete", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
