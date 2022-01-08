package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wpprof"
)

func main() {
	router := gin.Default()
	wpprof.Register(router)
	router.Run(":8080")
}
