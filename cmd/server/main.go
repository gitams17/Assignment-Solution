package main

import (
	"database/sql"
	"log"
	"os"
	"user-api/internal/logger"
	"user-api/internal/middleware"
	"user-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" 
)

const (
	dbDriver      = "postgres"
	serverAddress = ":8080"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		dbSource = "postgresql://root:secret@localhost:5432/userdb?sslmode=disable"
	}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		logger.Log.Fatal("Cannot connect to db")
	}

	if err := conn.Ping(); err != nil {
		logger.Log.Fatal("Cannot ping db")
	}

	app := fiber.New()

	middleware.SetupMiddleware(app)

	routes.SetupRoutes(app, conn)

	log.Fatal(app.Listen(serverAddress))
}