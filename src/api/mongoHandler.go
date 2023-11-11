package api

import "go.mongodb.org/mongo-driver/mongo"

type Handler struct {
	MongoClient *mongo.Client
}

func MongoHandler(mongoClient *mongo.Client) *Handler {
	return &Handler{
		MongoClient: mongoClient,
	}

}
