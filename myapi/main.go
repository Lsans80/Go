package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Entrada principal de la aplicación, donde se configura el enrutador y se inicia el servidor con un endpoint /api/hello que devuelve un mensaje JSON.
func main() {
	// Crear un enrutador
	router := mux.NewRouter()

	// Llamar a la función HelloHandler cuando se haga una petición GET a /api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")

	// Iniciar el servidor en el puerto 8080
	http.ListenAndServe(":8080", router)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Crear un mensaje JSON
	response := map[string]string{"message": "Hello Luisa, you are learning Go!"}

	// Establecer el tipo de contenido en la cabecera de la respuesta
	w.Header().Set("Content-Type", "application/json")

	// Codificar el mensaje JSON y enviarlo como respuesta
	json.NewEncoder(w).Encode(response)
}
