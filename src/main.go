package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/logmsg"

	_ "github.com/lib/pq"
)

type HandlerContext struct {
	DB *database.Queries
}

func main() {
	log.SetPrefix("[rss-aggregator]")
	log.SetFlags(log.Llongfile)

	log.Print(logmsg.Warning("test"))

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment, forgot .env file?")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment, forgot .env file?")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	hc1 := HandlerContext{DB: database.New(conn)}

	router := gin.Default()
	router.Use(corsMiddleware())

	v1 := router.Group("/v1")
	{
		v1.GET("/healthz", handleReady)
		v1.GET("/users", hc1.handleGetUser)
		v1.POST("/users", hc1.handleCreateUser)
	}

	if os.Getenv("HTTP_ONLY") == "true" {
		err = router.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = router.RunTLS(":"+port, "server.crt", "server.key")
		if err != nil {
			log.Print("Make sure to generate TLS certificate to enable HTTPS by running 'make generate-TLS'")
			log.Fatal(err)
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
