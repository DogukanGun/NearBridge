package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"networklistener/utils"
	"time"
)

/*
ConnectMongo function to connect mongo db
*/
func (dh *DatabaseHandler) ConnectMongo() *mongo.Client {
	envNames := []string{"MONGO_URI"}
	status, missingName, envMap := utils.InitializeENV(envNames, "mongo.env")
	if !status {
		fmt.Println(missingName + " was not initialized... Aborting!")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(envMap["MONGO_URI"]))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
