package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id          primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name        string             `bson:"name, omitempty"`
	Description string             `bson:"description, omitempty"`
}
