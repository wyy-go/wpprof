package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wyy-go/wpprof"
)

func main() {
	router := gin.Default()
	wpprof.Register(router,
		wpprof.WithPrefix("/admin"),
		wpprof.WithHandlers(func(c *gin.Context) {
			if c.Request.Header.Get("Authorization") != "foobar" {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
			c.Next()
		}),
	)
	router.Run(":8080")
}
