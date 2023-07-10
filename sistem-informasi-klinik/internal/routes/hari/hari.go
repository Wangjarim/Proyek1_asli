package hariRoutes

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	hariHandler "sistem-informasi-klinik/internal/handlers/hari"
)

func SetupHariRoutes(router fiber.Router) {
	user := router.Group("/hari")
	// Create a user
	user.Post("/", hariHandler.Create)
	// Read all users
	user.Get("/", hariHandler.Get)
	// // Read one user
	user.Get("/:id", hariHandler.GetHari)
	// // Update one user
	user.Put("/:id", hariHandler.Update)
	// // Delete one user
	user.Delete("/:id", hariHandler.Delete)
}
