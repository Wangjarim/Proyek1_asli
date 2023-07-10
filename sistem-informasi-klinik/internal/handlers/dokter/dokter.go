package dokterHandler

import (
	_ "github.com/gofiber/fiber/v2"
	fiber "github.com/gofiber/fiber/v2"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/internal/model"
)

func Get(c *fiber.Ctx) error {
	{
		var dokters []model.Dokter
		// Find all users in database
		result := database.DB.Find(&dokters)
		// Check for errors during query execution
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return list of users
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Dokter Berhasil Ditampilkan!",
			"data":    dokters,
		})
	}
}

func Create(c *fiber.Ctx) error {
	// Parse request body
	{
		var dokter model.Dokter
		if err := c.BodyParser(&dokter); err != nil {
			return err
		}
		// Insert new user into database
		result := database.DB.Create(&dokter)
		// Check for errors during insertion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Dokter Berhasil Ditambahkan!",
			"data":    dokter,
		})
	}
}
func GetDokter(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id")
		// Find user by id_user in database
		var dokter model.Dokter
		result := database.DB.First(&dokter, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Dokter Tidak Ditemukan!",
			})
		}
		// Check for errors during query
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return user
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
			"data":    dokter,
		})
	}
}
func Update(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id")
		// Find user by id_user in database
		var dokter model.Dokter
		result := database.DB.First(&dokter, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Dokter Tidak Ditemukan",
			})
		}
		// Parse request body
		var newDokter model.Dokter
		if err := c.BodyParser(&newDokter); err != nil {
			return err
		}
		// Update user in database
		result = database.DB.Model(&dokter).Updates(newDokter)
		// Check for errors during update
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Dokter Berhasil Diperbarui!",
			"data":    dokter,
		})
	}
}
func Delete(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	{
		id := c.Params("id")
		// Find user by id_user in database
		var dokter model.Dokter
		result := database.DB.First(&dokter, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Dokter Tidak Ditemukan",
			})
		}
		// Delete user from database
		result = database.DB.Delete(&dokter)
		// Check for errors during deletion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Dokter Berhasil Dihapus!",
			"data":    dokter,
		})
	}
}
