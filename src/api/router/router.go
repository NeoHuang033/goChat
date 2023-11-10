package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goChat/src/api/handler"
	"log"
	"net/http"
	"time"
)

var mongoClient *mongo.Client

func Register() *gin.Engine {
	r := gin.Default()
	setupMongoDB()
	//r.Use(CorsMiddleware())
	initUserRouter(r)
	r.Run(":8080")
	return r
}

func initUserRouter(router *gin.Engine) {
	//	userGroup := r.Group("/user")

	mongoHandler := handler.MongoHandler(mongoClient)
	router.POST("/login", mongoHandler.LoginUserHandler())
	router.POST("/register", mongoHandler.RegisterHandler())
	//userGroup.Use(CheckSessionId())

}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		var openCorsFlag = true
		if openCorsFlag {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
		}
		c.Next()
	}
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
