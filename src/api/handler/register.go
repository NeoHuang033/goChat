package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"goChat/src/domain"
	"net/http"
	"time"
)

func (h *Handler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.CreatedTime = time.Now()

		// user.PassWord = HashPassword(user.PassWord)

		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
		defer cancel()

		collection := h.MongoClient.Database("Neo-GoChat-Service").Collection("userInfo")
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}
}

/*func (h *Handler) RegisterHandlers(c *gin.Context) {
	user := domain.User{
		UserName:    c.GetString("name"),
		PassWord:    c.GetString("passWord"), //todo Hash the password here before inserting it into the database
		CreatedTime: time.Now(),
	}
	// Bind the JSON to user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

}*/
