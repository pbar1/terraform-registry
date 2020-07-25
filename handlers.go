package main

import (
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog/log"
)

// https://www.terraform.io/docs/internals/remote-service-discovery.html#discovery-process
func serviceDiscovery(c *fiber.Ctx) {
	log.Debug().Send()
	if err := c.JSON(fiber.Map{"modules.v1": "/v1/modules/"}); err != nil {
		log.Error().Err(err)
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}
}

// https://www.terraform.io/docs/registry/api.html#list-modules
func listModules(c *fiber.Ctx) {
	log.Debug().Send()
	namespace := c.Params("namespace")
	var query listModulesQuery
	if err := c.QueryParser(&query); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Interface("query", query).Send()

	// TODO

	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#search-modules
func searchModules(c *fiber.Ctx) {
	log.Debug().Send()
	var query searchModulesQuery
	if err := c.QueryParser(&query); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO

	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#list-available-versions-for-a-specific-module
func listVersionsForModule(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name, provider := c.Params("namespace"), c.Params("name"), c.Params("provider")
	if namespace == "" || name == "" || provider == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Str("name", name).Str("provider", provider)

	// TODO

	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#download-source-code-for-a-specific-module-version
func downloadSourceForModule(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name, provider, version := c.Params("namespace"), c.Params("name"), c.Params("provider"), c.Params("version")
	if namespace == "" || name == "" || provider == "" || version == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Str("name", name).Str("provider", provider).Str("version", version)

	if exists, err := backend.ModuleExists(namespace, name, provider, version); !exists {
		if err != nil {
			log.Error().Err(err)
			c.SendStatus(fiber.StatusInternalServerError)
			return
		}
		c.SendStatus(fiber.StatusNotFound)
		return
	}
	c.Set("X-Terraform-Get", "/"+moduleArchiveFilename)

	c.SendStatus(fiber.StatusNoContent)
}

// https://www.terraform.io/docs/registry/api.html#list-latest-version-of-module-for-all-providers
func listLatestModulesAllProviders(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name := c.Params("namespace"), c.Params("name")
	if namespace == "" || name == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	var query listLatestModulesAllProvidersQuery
	if err := c.QueryParser(&query); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Str("name", name).Interface("query", query)

	// TODO

	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#latest-version-for-a-specific-module-provider
func listLatestModuleForProvider(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name, provider := c.Params("namespace"), c.Params("name"), c.Params("provider")
	if namespace == "" || name == "" || provider == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Str("name", name).Str("provider", provider)

	// TODO

	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#get-a-specific-module
func getSpecificModule(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name, provider, version := c.Params("namespace"), c.Params("name"), c.Params("provider"), c.Params("version")
	if namespace == "" || name == "" || provider == "" || version == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Str("name", name).Str("provider", provider).Str("version", version)

	// TODO

	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#download-the-latest-version-of-a-module
func downloadLatestModule(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name, provider := c.Params("namespace"), c.Params("name"), c.Params("provider")
	if namespace == "" || name == "" || provider == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Str("name", name).Str("provider", provider)

	// TODO

	c.SendStatus(fiber.StatusNotImplemented)
}

// custom, for inspiration refer to: https://www.terraform.io/docs/cloud/api/modules.html
func uploadModule(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name, provider, version := c.Params("namespace"), c.Params("name"), c.Params("provider"), c.Params("version")
	if namespace == "" || name == "" || provider == "" || version == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	body := c.Body()
	log.Debug().Str("namespace", namespace).Str("name", name).Str("provider", provider).Str("version", version)

	if err := backend.PutModule(namespace, name, provider, version, []byte(body)); err != nil {
		log.Error().Err(err)
		c.SendStatus(fiber.StatusBadRequest)
		return
	}

	c.SendStatus(fiber.StatusCreated)
}

func downloadModule(c *fiber.Ctx) {
	log.Debug().Send()
	namespace, name, provider, version := c.Params("namespace"), c.Params("name"), c.Params("provider"), c.Params("version")
	if namespace == "" || name == "" || provider == "" || version == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	log.Debug().Str("namespace", namespace).Str("name", name).Str("provider", provider).Str("version", version)

	bindata, err := backend.GetModule(namespace, name, provider, version)
	if err != nil {
		log.Error().Err(err)
		c.SendStatus(fiber.StatusBadRequest)
		return
	}

	c.SendBytes(bindata)
	c.Type(fiber.MIMEOctetStream)
	c.SendStatus(fiber.StatusOK)
}
