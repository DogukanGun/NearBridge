package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

func (dh *DatabaseHandler) InsertData(client *mongo.Client, theData interface{}) (status bool, err error) {
	db := dh.targetDB
	collection := dh.targetTable

	coll := client.Database(db).Collection(collection)
	_, err = coll.InsertOne(context.TODO(), theData)

	if err != nil {
		fmt.Println(err)
		status = false
		return
	}
	status = true
	return
}
