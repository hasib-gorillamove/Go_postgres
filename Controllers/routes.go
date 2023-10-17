package controllers

import "github.com/gorilla/mux"

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/item", CreateItem).Methods("POST")
	r.HandleFunc("/item/{id:[0-9]+}", GetItem).Methods("GET")
	r.HandleFunc("/item/{id:[0-9]+}", UpdateItem).Methods("PUT")
	r.HandleFunc("/item/{id:[0-9]+}", DeleteItem).Methods("DELETE")
	r.HandleFunc("/item", GetAllItems).Methods("GET")
}
