package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello Luisa, you are learning Go!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
