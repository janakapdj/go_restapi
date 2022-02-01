package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

//Find all matching documents in collection
func Find(document string, filter interface{}, db *mongo.Database, limit int64) (*mongo.Cursor, error) {

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	cur, err := db.Collection(document).Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

//Find all matching documents in collection
func Aggregate(document string, pipeline interface{}, db *mongo.Database, limit int64) (*mongo.Cursor, error) {

	cur, err := db.Collection(document).Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

//InsertOne insert single document in to the collection
func InsertOne(document string, entity interface{}, db *mongo.Database) (interface{}, error) {
	insertResult, err := db.Collection(document).InsertOne(context.TODO(), entity)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}


//UpdateOne update single document in collection
func UpdateOne(document string, filter interface{}, entity interface{}, db *mongo.Database) (int64, error) {
	res, err := db.Collection(document).UpdateOne(context.TODO(), filter, entity)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

//Delete delete single document in collection
func Delete(document string, filter interface{}, db *mongo.Database) (int64, error) {

	res, err := db.Collection(document).DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
