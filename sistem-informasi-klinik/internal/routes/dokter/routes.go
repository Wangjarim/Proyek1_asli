package dokterRoutes

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	dokterHandler "sistem-informasi-klinik/internal/handlers/dokter"
)

func SetupDokterRoutes(router fiber.Router) {
	user := router.Group("/dokter")
	// Create a user
	user.Post("/", dokterHandler.Create)
	// Read all users
	user.Get("/", dokterHandler.Get)
	// // Read one user
	user.Get("/:id", dokterHandler.GetDokter)
	// // Update one user
	user.Put("/:id", dokterHandler.Update)
	// // Delete one user
	user.Delete("/:id", dokterHandler.Delete)
}
