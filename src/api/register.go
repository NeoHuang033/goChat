package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"goChat/src/domain"
	"net/http"
	"time"
)

type Handler struct {
	MongoClient *mongo.Client
}

// NewHandler creates a new Handler with the given dependencies
func NewHandler(mongoClient *mongo.Client) *Handler {
	return &Handler{
		MongoClient: mongoClient,
	}
}

func (h *Handler) RegisterHandlers(c *gin.Context) {
	user := domain.User{
		UserName:    c.GetString("name"),
		PassWord:    c.GetString("passWord"), // 这里应该是密码的哈希值
		CreatedTime: time.Now(),
	}
	// Bind the JSON to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password here before inserting it into the database
	// user.PassWord = HashPassword(user.PassWord)

	// Insert the user into the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := h.MongoClient.Database("yourDatabaseName").Collection("userInfo")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func (h *Handler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这个闭包内部，你现在可以使用 c *gin.Context
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.CreatedTime = time.Now()
		// 假设HashPassword是一个哈希密码的函数
		// user.PassWord = HashPassword(user.PassWord)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		collection := h.MongoClient.Database("Neo-Gochat-Service").Collection("userInfo")
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}
}
