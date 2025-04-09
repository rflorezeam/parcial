package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rflorezeam/parcial/config"
	controller "github.com/rflorezeam/parcial/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error cargando .env")
	}

	// Conectar a MongoDB
	config.ConectarDB()

	// Crear router
	router := mux.NewRouter()

	// Aquí luego registraremos los endpoints
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Microservicio de Libros")
	}).Methods("GET")

	// Correr servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Rutas CRUD
	router.HandleFunc("/libros", controller.CrearLibro).Methods("POST")
	router.HandleFunc("/libros", controller.ObtenerLibros).Methods("GET")
	router.HandleFunc("/libros/{id}", controller.ObtenerLibroPorID).Methods("GET")
	router.HandleFunc("/libros/{id}", controller.ActualizarLibro).Methods("PUT")
	router.HandleFunc("/libros/{id}", controller.EliminarLibro).Methods("DELETE")

	fmt.Println("🟢 Servidor corriendo en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
