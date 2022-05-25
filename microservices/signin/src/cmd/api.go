package cmd

import (
	"banktest_signin/src/service"
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

	app.Post("/signin", service.SignIn)
	app.Post("/create", service.Create)
	app.Post("/signout", service.SignOut)
	app.Post("/validate", service.Validate)

	err := app.Listen(":28080")
	if err != nil {
		panic(err)
	}
}
