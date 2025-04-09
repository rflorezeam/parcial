// repository/libro_repository.go
package repository

import (
	"context"
	"time"

	"github.com/rflorezeam/parcial/config"
	model "github.com/rflorezeam/parcial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var timeout = 10 * time.Second

func CrearLibro(libro model.Libro) (*model.Libro, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var collection = config.GetCollection("librosdb")

	libro.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, libro)
	return &libro, err
}

func ObtenerLibros() ([]model.Libro, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var collection = config.GetCollection("librosdb")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var libros []model.Libro
	if err = cursor.All(ctx, &libros); err != nil {
		return nil, err
	}

	return libros, nil
}

func ObtenerLibroPorID(id string) (*model.Libro, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var collection = config.GetCollection("librosdb")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var libro model.Libro
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&libro)
	return &libro, err
}

func ActualizarLibro(id string, libro model.Libro) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var collection = config.GetCollection("librosdb")

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": libro},
	)
	return err
}

func EliminarLibro(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var collection = config.GetCollection("librosdb")

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
