package main

import (
	"goChat/src/api"
	"goChat/src/connect"
	//"fmt"
)

func main() {
	/*secretKey, err := generateSecretKey()
	if err != nil {
		panic(err)
	}
	fmt.Println("Secret Key:", secretKey)*/

	go connect.New().Run()
	api.New().Run()

}
