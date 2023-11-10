package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson"
	"goChat/src/domain"
	"log"
	"net/http"
	"time"
)

const SecretKey = "15f5a4e7cfb1ec334f67433af0a4f86e563066945f7337314bd104a373fee500"

func (h *Handler) LoginUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		collection := h.MongoClient.Database("Neo-GoChat-Service").Collection("userInfo")
		err := collection.FindOne(ctx, bson.M{"username": user.UserName, "password": user.PassWord}).Decode(&user)
		if err != nil {
			log.Printf("FindOne error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't find user"})
			return
		}
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = user.Id
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		tokenString, err := token.SignedString([]byte(SecretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": tokenString})
	}
}
