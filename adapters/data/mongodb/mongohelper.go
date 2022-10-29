package mongodb

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/gitdoc/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/gitdoc/application"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoHelper struct {
	coll *mongo.Collection
}

func (mh mongoHelper) Find(ctx context.Context) ([]dao.DocumentDAO, error) {
	var documentDAOs = make([]dao.DocumentDAO, 0)
	cur, err := mh.coll.Find(ctx, bson.M{})
	if err != nil {
		log.Error().Err(err).Msgf("Error getting documents")
		return nil, err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, &documentDAOs)
	return documentDAOs, err
}

func (mh mongoHelper) InsertOne(ctx context.Context, document interface{}) (string, error) {
	result, err := mh.coll.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", result.InsertedID), nil
}
func (mh mongoHelper) FindOne(ctx context.Context, id string) (dao.DocumentDAO, error) {
	var documentDAO dao.DocumentDAO
	err := mh.coll.FindOne(ctx, bson.M{"uuid": id}).Decode(&documentDAO)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting document")
		return dao.DocumentDAO{}, &application.ErrorCannotFinddocument{ID: id}
	}
	return documentDAO, nil
}

func (mh mongoHelper) UpdateOne(ctx context.Context, id string, update interface{}) (int, error) {
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	result, err := mh.coll.UpdateOne(ctx, bson.M{"uuid": id}, update, &updateOpts)
	return int(result.ModifiedCount), err
}

func (mh mongoHelper) DeleteOne(ctx context.Context, id string) (int, error) {
	result, err := mh.coll.DeleteOne(ctx, bson.M{"uuid": id})
	return int(result.DeletedCount), err
}
