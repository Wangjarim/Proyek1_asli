package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	belajarRoutes "sistem-informasi-klinik/internal/routes/belajar"
	dokterRoutes "sistem-informasi-klinik/internal/routes/dokter"
	hariRoutes "sistem-informasi-klinik/internal/routes/hari"
	jadwal_dokterRoutes "sistem-informasi-klinik/internal/routes/jadwal_dokter"
	jamRoutes "sistem-informasi-klinik/internal/routes/jam"
	ruanganRoutes "sistem-informasi-klinik/internal/routes/ruangan"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	// Setup the Node Routes
	belajarRoutes.SetupBelajarRoutes(api)
	dokterRoutes.SetupDokterRoutes(api)
	hariRoutes.SetupHariRoutes(api)
	jamRoutes.SetupJamRoutes(api)
	jadwal_dokterRoutes.SetupJadwal_dokterRoutes(api)
	ruanganRoutes.SetupRuanganRoutes(api)
}
