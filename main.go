package main

import (
	"Go_postgres/controllers"
	"Go_postgres/database"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()
	controllers.SetupRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
