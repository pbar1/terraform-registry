package main

import (
	"github.com/dre1080/fiberlog"
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog/log"
)

func NewServer() *fiber.App {
	app := fiber.New(&fiber.Settings{
		StrictRouting:         true,
		DisableStartupMessage: true,
	})

	app.Use(fiberlog.New(fiberlog.Config{
		Logger: &log.Logger,
	}))

	app.Get("/.well-known/terraform.json", serviceDiscovery)

	v1modules := app.Group("/v1/modules", listModules)
	v1modules.Get("/search", searchModules)
	v1modules.Get("/:namespace", listModules)
	v1modules.Get("/:namespace/:name", listLatestModulesAllProviders)
	v1modules.Get("/:namespace/:name/:provider", listLatestModuleForProvider)
	v1modules.Get("/:namespace/:name/:provider/download", downloadLatestModule)
	v1modules.Get("/:namespace/:name/:provider/versions", listVersionsForModule)
	v1modules.Get("/:namespace/:name/:provider/:version", getSpecificModule)
	v1modules.Get("/:namespace/:name/:provider/:version/download", downloadSourceForModule)

	return app
}
