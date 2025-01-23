package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/shelllbyyyyy/belajar-api-go/apps/auth"
	"github.com/shelllbyyyyy/belajar-api-go/external/database"
	"github.com/shelllbyyyyy/belajar-api-go/internal/environtment"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	pwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    
    err = godotenv.Load(filepath.Join(pwd, "internal", "environtment" , ".env"))
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	cfg, err := environtment.LoadConfig();
	if err != nil {
		panic(err)
	}
	
	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	auth.UserRoute(app, db)
	auth.AuthRoute(app, db)

	app.Listen(":3000")
}