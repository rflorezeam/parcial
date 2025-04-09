// controllers/libro_controller.go
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/parcial/models"
	service "github.com/rflorezeam/parcial/services"
)

func CrearLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	creado, err := service.CrearLibro(libro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(creado)
}

func ObtenerLibros(w http.ResponseWriter, r *http.Request) {
	libros, err := service.ObtenerLibros()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(libros)
}

func ObtenerLibroPorID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	libro, err := service.ObtenerLibroPorID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(libro)
}

func ActualizarLibro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var libro models.Libro
	if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.ActualizarLibro(id, libro); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func EliminarLibro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := service.EliminarLibro(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
