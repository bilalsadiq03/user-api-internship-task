package main

import (
	"database/sql"
	"log"


	
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/bilalsadiq03/user-api-internship-task/config"
	"github.com/bilalsadiq03/user-api-internship-task/internal/middleware"
	"github.com/bilalsadiq03/user-api-internship-task/internal/handler"
	"github.com/bilalsadiq03/user-api-internship-task/internal/logger"
	"github.com/bilalsadiq03/user-api-internship-task/internal/repository"
	"github.com/bilalsadiq03/user-api-internship-task/internal/routes"
)
	


func main() {

	cfg := config.Load()

	logg := logger.New()
	defer logg.Sync()

	db, err := sql.Open("postgres", cfg.DBurl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	
	app := fiber.New()

	// Middlewares
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(logg))


	// Dependencies
	repo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(repo, logg)


	// Routes
	routes.Register(app, userHandler)


	log.Fatal(app.Listen(":" + cfg.Port))
}