package main

import (
	"crypto/rand"
	"encoding/hex"
	//"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"goChat/src/api"
	"goChat/src/connect"
)

var mongoClient *mongo.Client

func main() {
	/*secretKey, err := generateSecretKey()
	if err != nil {
		panic(err)
	}
	fmt.Println("Secret Key:", secretKey)*/
	connect.New().Run()
	api.New().Run()

}

func generateSecretKey() (string, error) {
	bytes := make([]byte, 32) // 生成 256 位密钥
	if _, err := rand.Read(bytes); err != nil {
		return "", err // 读取随机数失败
	}
	return hex.EncodeToString(bytes), nil // 将字节序列转换成十六进制编码的字符串
}
