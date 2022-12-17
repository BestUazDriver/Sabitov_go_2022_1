package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(mongoCollection *mongo.Collection) *PostRepository {
	return &PostRepository{
		collection: mongoCollection,
	}
}

//
//func GetAll(ctx *context.Context) ([]*PostRepository, error) {
//
//}
//
//func GetById(ctx context.Context, id string) (*core.User, error) {
//	objectId, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return nil, err
//	}
//
//}
