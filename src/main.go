package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/handler"
	"github.com/pewpewnor/rss-aggregator/src/logmsg"
	"github.com/pewpewnor/rss-aggregator/src/utils"

	_ "github.com/lib/pq"
)

func main() {
	log.SetPrefix("[rss-aggregator]")
	log.SetFlags(log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(logmsg.Error(err))
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(logmsg.Error(
			"PORT is not found in the environment, forgot .env file?"))
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal(logmsg.Error(
			"DB_URL is not found in the environment, forgot .env file?"))
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(logmsg.Error("Can't connect to database"))
	}

	db := database.New(conn)

	go utils.StartScraping(db, 10, time.Minute)

	hc1 := handler.HandlerContext{DB: db}

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
		v1WithAuth.DELETE("/subscribes/:subscribeID", hc1.HandleDeleteSubscribe)
	}

	log.Println(logmsg.Info("server trying to start at port ", port))

	if os.Getenv("HTTP_ONLY") == "true" {
		log.Println(logmsg.Info("server set to only support HTTP"))

		err = router.Run()
		if err != nil {
			log.Fatal(logmsg.Error(err))
		}
	} else {
		log.Println(logmsg.Info("server set to support HTTPS and HTTP"))

		err = router.RunTLS(":"+port, "server.crt", "server.key")
		if err != nil {
			log.Println(logmsg.Error(
				"Make sure to generate TLS certificate to enable HTTPS by running 'make generate-TLS'"))
			log.Fatal(logmsg.Error(err))
		}
	}

	log.Println(logmsg.Success("server sucessfully exited from port ", port))
}
