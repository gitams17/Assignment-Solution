package routes

import (
	"database/sql"
	"user-api/db/sqlc"
	"user-api/internal/handler"
	"user-api/internal/repository"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, dbConn *sql.DB) {
	q := db.New(dbConn)
	
	repo := repository.NewRepository(q)
	
	svc := service.NewUserService(repo.(*repository.SQLStore))
	
	h := handler.NewUserHandler(svc)

	api := app.Group("/users")

	api.Post("/", h.CreateUser)      // Create
	api.Get("/", h.ListUsers)        // List All
	api.Get("/:id", h.GetUser)       // Get One
	api.Put("/:id", h.UpdateUser)    // Update
	api.Delete("/:id", h.DeleteUser) // Delete
}