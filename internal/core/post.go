package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id      primitive.ObjectID
	Likes   int
	Owner   *User
	Content string
}
