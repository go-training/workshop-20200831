package main

import (
	"flag"
	"log"
	"net/http"

	"gin-http-server/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	var server = flag.Bool("server", false, "enable server")
	var ping = flag.Bool("ping", false, "ping server")
	flag.Parse()

	config.Load()

	r := gin.Default()
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user": id,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		data := &user{
			Name:  "foo",
			Email: "foo@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
		})
	})

	r.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		data := &user{
			Name:  "foo",
			Email: "foo@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
			"id":   id,
		})
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		data := &user{
			Name:  "foo",
			Email: "foo@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
			"id":   id,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	if *server {
		r.Run(":" + config.Setting.Port) // listen on 0.0.0.0:8081
	}

	if *ping {
		resp, err := http.Get("http://localhost:" + config.Setting.Port + "/healthz")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Fatal("can't connect to server")
		}

		log.Println("connected to the server")
	}
}
