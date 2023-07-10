package ruanganRoutes

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	ruanganHandler "sistem-informasi-klinik/internal/handlers/ruangan"
)

func SetupRuanganRoutes(router fiber.Router) {
	user := router.Group("/ruangan")
	// Create a user
	user.Post("/", ruanganHandler.Create)
	// Read all users
	user.Get("/", ruanganHandler.Get)
	// // Read one user
	user.Get("/:id", ruanganHandler.GetRuangan)
	// // Update one user
	user.Put("/:id", ruanganHandler.Update)
	// // Delete one user
	user.Delete("/:id", ruanganHandler.Delete)
}
