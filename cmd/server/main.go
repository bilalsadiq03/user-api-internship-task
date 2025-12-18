package main

import (
	"database/sql"
	"log"


	
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"github.com/bilalsadiq03/user-api-internship-task/internal/middleware"
	"github.com/bilalsadiq03/user-api-internship-task/internal/handler"
	"github.com/bilalsadiq03/user-api-internship-task/internal/logger"
	"github.com/bilalsadiq03/user-api-internship-task/internal/repository"
	"github.com/bilalsadiq03/user-api-internship-task/internal/routes"
)
	


func main() {

	db, err := sql.Open("postgres", "postgres://postgres:9528296572@localhost:5432/user_api?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.New()
	
	app := fiber.New()

	// Middlewares
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(logg))

	repo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(repo, logg)

	routes.Register(app, userHandler)


	log.Fatal(app.Listen(":8080"))
}