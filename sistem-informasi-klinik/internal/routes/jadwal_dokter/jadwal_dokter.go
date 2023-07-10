package jadwal_dokterRoutes

import (
	"github.com/gofiber/fiber/v2"
	jadwal_dokterHandler "sistem-informasi-klinik/internal/handlers/jadwal_dokter"
)

func SetupJadwal_dokterRoutes(router fiber.Router) {
	user := router.Group("/jadwal_dokter")
	// Create a user
	user.Post("/", jadwal_dokterHandler.Create)
	// Read all users
	user.Get("/", jadwal_dokterHandler.Get)
	// // Read one user
	user.Get("/:id", jadwal_dokterHandler.GetJadwal)
	// // Update one user
	user.Put("/:id", jadwal_dokterHandler.Update)
	// // Delete one user
	user.Delete("/:id", jadwal_dokterHandler.Delete)
}
