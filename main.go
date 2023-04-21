package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/passwords/:length", func(c *gin.Context) {
		length, err := strconv.Atoi(c.Param("length"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password length"})
			return
		}

		if length < 8 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password length must be at least 8"})
			return
		}

		password := generatePassword(length)
		c.JSON(http.StatusOK, gin.H{"password": password})
	})

	router.Run(":8080")

}

func generatePassword(length int) string {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		panic(err)
	}

	password := base64.URLEncoding.EncodeToString(buffer)[:length]
	return password
}
