//  Neutron Knock - Main
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/helmet/v2"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// * Check dev mode. Branching paths for retrieving necessary credentials
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "DEVELOPMENT"
		err := godotenv.Load()
		if err != nil {
			log.Fatal("ENVIRONMENT COULD NOT BE LOADED")
		}
	}

	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		AppName:      "Neutron Knock v0.1a",
		ServerHeader: "Neutron v0.1a",
	})

	app.Use(cache.New())
	app.Use(helmet.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Neutron Knock Server Metrics"}))

	app.Listen(":" + port)

	scheduler := gocron.NewScheduler(time.Local)

	app.Get("/", func(ctx *fiber.Ctx) error {
		scheduler.Every(5).Seconds().Do(func() {
			fmt.Println("New Task Scheduled")
		})
		return ctx.SendString("New Task Scheduled")
	})

	scheduler.Every(1).Seconds().Do(func() {

		log.Println("Hello World!")
	})

	scheduler.StartAsync()
	log.Println(app.Listen(":" + port))

}
