package main

import (
	"crypto/rand"
	"encoding/hex"
	"goChat/src/api"

	"goChat/src/connect"

	//"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client

func main() {
	/*secretKey, err := generateSecretKey()
	if err != nil {
		panic(err)
	}
	fmt.Println("Secret Key:", secretKey)*/

	go connect.New().Run()
	api.New().Run()

}

func generateSecretKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err // 读取随机数失败
	}
	return hex.EncodeToString(bytes), nil
}
