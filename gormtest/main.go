package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"test.com/gormtest/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
