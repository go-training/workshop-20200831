package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"gin-http-server/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func router() http.Handler {
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
		c.AbortWithStatus(200)
	})

	return r
}

func main() {
	config.Load()
	var server = flag.Bool("server", false, "enable server")
	var ping = flag.Bool("ping", false, "ping server")
	flag.Parse()

	hander := router()

	s := http.Server{
		Addr:         ":" + config.Setting.Server.Port,
		Handler:      hander,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if *server {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}

	if *ping {
		resp, err := http.Get("http://localhost:" + config.Setting.Server.Port + "/healthz")
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
