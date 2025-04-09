package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rflorezeam/parcial/config"
	"github.com/rflorezeam/parcial/models"
	service "github.com/rflorezeam/parcial/services"

	"github.com/stretchr/testify/assert"
)

func TestCrearLibro(t *testing.T) {
	// Cargar .env
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("❌ No se pudo cargar el archivo .env")
	}

	// Verificar que la URI esté bien formada
	if os.Getenv("MONGO_URI") == "" {
		t.Fatal("❌ MONGO_URI está vacío o no cargado")
	}
	config.ConectarDB()
	libro := models.Libro{
		Titulo:          "Prueba",
		Autor:           "Pepito",
		Editorial:       "Editorial",
		AnioPublicacion: 2000,
		Disponible:      false,
	}

	creado, err := service.CrearLibro(libro)
	assert.Nil(t, err)
	assert.Equal(t, libro.Titulo, creado.Titulo)

	// Obtener por ID
	encontrado, err := service.ObtenerLibroPorID(creado.ID.Hex())
	assert.Nil(t, err)
	assert.Equal(t, creado.ID, encontrado.ID)
}

func TestObtenerLibros(t *testing.T) {
	Libros, err := service.ObtenerLibros()

	assert.NoError(t, err)
	assert.NotNil(t, Libros)
}

func TestActualizarLibro(t *testing.T) {
	// Crear una Libro para actualizar
	Libro := models.Libro{
		Titulo:          "Prueba",
		Autor:           "Pepito",
		Editorial:       "Editorial",
		AnioPublicacion: 2000,
		Disponible:      false,
	}
	creado, err := service.CrearLibro(Libro)
	assert.NoError(t, err)

	// Datos actualizados
	libroActualizado := models.Libro{
		Titulo:          "Test",
		Autor:           "Arturo",
		Editorial:       "Editorial2",
		AnioPublicacion: 1000,
		Disponible:      true,
	}

	err = service.ActualizarLibro(creado.ID.Hex(), libroActualizado)
	assert.NoError(t, err)
}

func TestEliminarLibro(t *testing.T) {
	// Crear una tarea para luego eliminar
	libro := models.Libro{
		Titulo:          "Prueba",
		Autor:           "Pepito",
		Editorial:       "Editorial",
		AnioPublicacion: 2000,
		Disponible:      false,
	}
	creado, err := service.CrearLibro(libro)
	assert.NoError(t, err)

	// Eliminar la libro
	err = service.EliminarLibro(creado.ID.Hex())
	assert.NoError(t, err)
}
