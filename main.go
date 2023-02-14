//  Neutron Knock - Main
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/helmet/v2"

	"github.com/gofiber/fiber/v2"
	endpoints "neutron.money/knock/endpoints"
	scheduler "neutron.money/knock/scheduler"
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

	schedulerObject := scheduler.GetScheduler()

	neutronAPIS := app.Group("/")
	endpoints.AffixJobsRoutes(&neutronAPIS)

	schedulerObject.StartAsync()
	app.Listen(":" + port)

}
