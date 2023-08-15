package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/handler"

	_ "github.com/lib/pq"
)

func main() {
	log.SetPrefix("[rss-aggregator]")
	log.SetFlags(log.Llongfile)

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

	hc1 := handler.HandlerContext{DB: database.New(conn)}

	router := gin.Default()
	router.Use(handler.CORSMiddleware())
	router.Use(hc1.AliveMiddleware())

	v1NoAuth := router.Group("/v1")
	{
		v1NoAuth.GET("/healthz", handler.HandleReady)

		v1NoAuth.POST("/users", hc1.HandleCreateUser)
	}

	v1WithAuth := router.Group("/v1")
	v1WithAuth.Use(hc1.AuthMiddleware())
	{
		v1WithAuth.GET("/users", hc1.HandleGetUser)

		v1WithAuth.POST("/feeds", hc1.HandleCreateFeed)

		v1WithAuth.GET("/subscribes", hc1.HandleGetSubscribe)
		v1WithAuth.POST("/subscribes", hc1.HandleCreateSubscribe)
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
