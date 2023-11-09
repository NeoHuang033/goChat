package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID
	UserName    string
	PassWord    string
	CreatedTime time.Time
}
