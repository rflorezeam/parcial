// service/libro_service.go
package service

import (
	"github.com/rflorezeam/parcial/models"
	repository "github.com/rflorezeam/parcial/repositories"
)

func CrearLibro(libro models.Libro) (*models.Libro, error) {
	return repository.CrearLibro(libro)
}

func ObtenerLibros() ([]models.Libro, error) {
	return repository.ObtenerLibros()
}

func ObtenerLibroPorID(id string) (*models.Libro, error) {
	return repository.ObtenerLibroPorID(id)
}

func ActualizarLibro(id string, libro models.Libro) error {
	return repository.ActualizarLibro(id, libro)
}

func EliminarLibro(id string) error {
	return repository.EliminarLibro(id)
}
