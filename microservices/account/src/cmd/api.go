package cmd

import (
	"banktest_account/src/service"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"
)

func StartHttp(ctx context.Context) {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	app.Get("/liveness", service.Liveness)
	app.Get("/health", service.Health)

	app.Post("/create", service.Create)
	app.Get("/account", service.BankAccount)
	app.Post("/balance", service.UpdateBalance)
	app.Post("/status", service.Status)
	err := app.Listen(":28081")
	if err != nil {
		panic(err)
	}
}
