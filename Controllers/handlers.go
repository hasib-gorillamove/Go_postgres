package controllers

import (
	"Go_postgres/database"
	"Go_postgres/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// implementation of createItem function
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid json"+err.Error(), http.StatusBadRequest)
		return
	}
	id, err := database.CreateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	item.ID = id

	response := map[string]string{"message ": "Cerate done"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
	w.Write(jsonResponse)

}

// implementaion of GetyAllItems function
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	item, err := database.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "getting done"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
	w.Write(jsonResponse)
}

// implementation of GetGetelementbyID function
func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := database.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.UpdateItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message ": "Update Done"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
	w.Write(jsonResponse)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DeleteItem(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Delete successful"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
