package jadwal_dokterHandler

import (
	"errors"
	_ "github.com/gofiber/fiber/v2"
	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/internal/model"
	//"gorm.io/gorm"
)

//type JdHandler struct {
//	JadwalDokter *jadwaldokter.JadwalDokterTable
//}

func Get(c *fiber.Ctx) error {
	var jadwals []model.JadwalDokter
	// Find all users in database
	result := database.DB.Preload("Dokter").Preload("Hari").Preload("Jam").Preload("Ruangan").Find(&jadwals)

	// Check for errors during query execution
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return list of users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Jadwal Berhasil Ditampilkan!",
		"data":    jadwals,
	})

	//var jadwal []model.JadwalDokter
	//
	//// Find all users in database and join with DataPribadi
	//result := database.DB.
	//	Select("jadwal_dokter.id_dokter, dokter.nama, dokter.keahlian").
	//	Joins("JOIN dokter ON jadwal_dokter.id_dokter = dokter.id").
	//	Find(&jadwal)
	//
	//// Check for errors during query execution
	//if result.Error != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"message": result.Error.Error(),
	//	})
}

// Return list of users
//return c.Status(fiber.StatusOK).JSON(fiber.Map{
//	"message": "Data User Berhasil Ditampilkan!",
//	"data":    jadwals,
//})

func Create(c *fiber.Ctx) error {
	// Parse request body
	{
		var jadwal model.JadwalDokter
		if err := c.BodyParser(&jadwal); err != nil {
			return err
		}
		// Insert new user into database
		result := database.DB.Create(&jadwal)
		// Check for errors during insertion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Jadwal Berhasil Ditambahkan!",
			"data":    jadwal,
		})
	}
}
func GetJadwal(c *fiber.Ctx) error {
	// Ambil ID dari parameter URL
	id := c.Params("id")

	var jadwal model.JadwalDokter
	// Cari jadwal dokter berdasarkan ID
	result := database.DB.Preload("Dokter").Preload("Hari").Preload("Jam").Preload("Ruangan").First(&jadwal, id)

	// Periksa apakah jadwal dokter ditemukan atau tidak
	if result.Error != nil {
		// Jika tidak ditemukan, kirim respons dengan status Not Found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Jadwal Dokter tidak ditemukan.",
			})
		}
		// Jika terjadi kesalahan lain saat query, kirim respons dengan status Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Jika jadwal dokter ditemukan, kirim respons dengan data jadwal dokter tersebut
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Jadwal Dokter berhasil ditemukan!",
		"data":    jadwal,
	})
}
func Update(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id")
		// Find user by id_user in database
		var jadwal model.JadwalDokter
		result := database.DB.First(&jadwal, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Jadwal Tidak Ditemukan",
			})
		}
		// Parse request body
		var newJadwal model.JadwalDokter
		if err := c.BodyParser(&newJadwal); err != nil {
			return err
		}
		// Update user in database
		result = database.DB.Model(&jadwal).Updates(newJadwal)
		// Check for errors during update
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Jadwal Berhasil Diperbarui!",
			"data":    jadwal,
		})
	}
}
func Delete(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	{
		id := c.Params("id")
		// Find user by id_user in database
		var jadwal model.JadwalDokter
		result := database.DB.First(&jadwal, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Jadwal Tidak Ditemukan",
			})
		}
		// Delete user from database
		result = database.DB.Delete(&jadwal)
		// Check for errors during deletion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Jadwal Berhasil Dihapus!",
			"data":    jadwal,
		})
	}
}
