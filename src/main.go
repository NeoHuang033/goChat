package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goChat/src/api"
)

var mongoClient *mongo.Client

func main() {
	setupMongoDB()

	// 创建Gin引擎
	router := gin.Default()
	handler := api.NewHandler(mongoClient)

	// 路由和处理器
	router.POST("/register", handler.RegisterHandler())
	//router.POST("/login", loginHandler)

	// 启动Gin服务器
	router.Run(":8080")
}

func setupMongoDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// 连接到MongoDB
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://neo:123456admin@localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
}
