package main

import (
	"crud-golang/db"
	"crud-golang/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// // Panjang secret key yang diinginkan (dalam byte)
	// keyLength := 32

	// // Buat slice untuk menampung nilai acak
	// key := make([]byte, keyLength)

	// // Baca nilai acak dari crypto/rand
	// _, err := rand.Read(key)
	// if err != nil {
	// 	fmt.Println("Failed to generate random key:", err)
	// 	return
	// }

	// // Encode nilai acak dalam base64
	// keyBase64 := base64.StdEncoding.EncodeToString(key)

	// fmt.Println("Random secret key (base64):", keyBase64)
	db.InitDB()

	r := gin.Default()
	routers.SetupRouter(r)

	r.Run() // default port :8080
}
