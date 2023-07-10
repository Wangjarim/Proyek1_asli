package jamRoutes

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	jamHandler "sistem-informasi-klinik/internal/handlers/jam"
)

func SetupJamRoutes(router fiber.Router) {
	user := router.Group("/jam")
	// Create a user
	user.Post("/", jamHandler.Create)
	// Read all users
	user.Get("/", jamHandler.Get)
	// // Read one user
	user.Get("/:id", jamHandler.GetJam)
	// // Update one user
	user.Put("/:id", jamHandler.Update)
	// // Delete one user
	user.Delete("/:id", jamHandler.Delete)
}
