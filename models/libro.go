// model/libro.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Libro struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo          string             `bson:"titulo" json:"titulo"`
	Autor           string             `bson:"autor" json:"autor"`
	Editorial       string             `bson:"editorial" json:"editorial"`
	AnioPublicacion int                `bson:"anio_publicacion" json:"anio_publicacion"`
	Disponible      bool               `bson:"disponible" json:"disponible"`
}
