package main

import "github.com/gofiber/fiber"

// https://www.terraform.io/docs/internals/remote-service-discovery.html#discovery-process
func serviceDiscovery(c *fiber.Ctx) {
	c.Status(fiber.StatusOK).Send(c.JSON(map[string]string{
		"modules.v1": "/v1/modules/",
	}))
}

// https://www.terraform.io/docs/registry/api.html#list-modules
func listModules(c *fiber.Ctx) {
	// namespace := c.Params("namespace")
	var query listModulesQuery
	if err := c.QueryParser(&query); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#search-modules
func searchModules(c *fiber.Ctx) {
	var query searchModulesQuery
	if err := c.QueryParser(&query); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#list-available-versions-for-a-specific-module
func listVersionsForModule(c *fiber.Ctx) {
	namespace, name, provider := c.Params("namespace"), c.Params("name"), c.Params("provider")
	if namespace == "" || name == "" || provider == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#download-source-code-for-a-specific-module-version
func downloadSourceForModule(c *fiber.Ctx) {
	namespace, name, provider, version := c.Params("namespace"), c.Params("name"), c.Params("provider"), c.Params("version")
	if namespace == "" || name == "" || provider == "" || version == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.Set("X-Terraform-Get", "TODO")
	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#list-latest-version-of-module-for-all-providers
func listLatestModulesAllProviders(c *fiber.Ctx) {
	namespace, name := c.Params("namespace"), c.Params("name")
	if namespace == "" || name == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	var query listLatestModulesAllProvidersQuery
	if err := c.QueryParser(&query); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#latest-version-for-a-specific-module-provider
func listLatestModuleForProvider(c *fiber.Ctx) {
	namespace, name, provider := c.Params("namespace"), c.Params("name"), c.Params("provider")
	if namespace == "" || name == "" || provider == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#get-a-specific-module
func getSpecificModule(c *fiber.Ctx) {
	namespace, name, provider, version := c.Params("namespace"), c.Params("name"), c.Params("provider"), c.Params("version")
	if namespace == "" || name == "" || provider == "" || version == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.SendStatus(fiber.StatusNotImplemented)
}

// https://www.terraform.io/docs/registry/api.html#download-the-latest-version-of-a-module
func downloadLatestModule(c *fiber.Ctx) {
	namespace, name, provider := c.Params("namespace"), c.Params("name"), c.Params("provider")
	if namespace == "" || name == "" || provider == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}
	c.SendStatus(fiber.StatusNotImplemented)
}
