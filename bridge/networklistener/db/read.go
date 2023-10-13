package db

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (dh *DatabaseHandler) GetData(client *mongo.Client, collection string, filter interface{}, unmarshalTarget interface{}) (err error) {

	cursor, err := client.Database(dh.targetDB).Collection(collection).Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	isDataExists := cursor.Next(context.Background())
	if isDataExists {
		var decodeArr []bson.M
		err = cursor.All(context.Background(), &decodeArr)
		marshalledData, _ := json.Marshal(decodeArr)
		err = json.Unmarshal(marshalledData, unmarshalTarget)
	} else {
		fmt.Println("Data is not exists")
	}
	return
}
