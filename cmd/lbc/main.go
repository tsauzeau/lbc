package main

import (
	"fmt"
	"net/http"

	"github.com/tsauzeau/lbc/cmd/lbc/config"
	"github.com/tsauzeau/lbc/cmd/lbc/controllers"
	"github.com/tsauzeau/lbc/cmd/lbc/db"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	r := gin.Default()

	r.Use(CORSMiddleware())

	db.Init(config.Config.RedisHost)

	v1 := r.Group("/v1")
	{
		/*** START Fizzbuzz ***/
		fizzbuzz := new(controllers.FizzbuzzController)

		v1.GET("/fizzbuzz", fizzbuzz.Get)
		v1.GET("/stat", fizzbuzz.Stat)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Welcome to fizzbuzz api, go to /v1/fizzbuzz"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"Message": "Page Not Found"})
	})

	r.Run(config.Config.APIPort)
}
