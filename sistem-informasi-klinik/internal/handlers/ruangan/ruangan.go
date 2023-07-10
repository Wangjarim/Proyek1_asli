package ruanganHandler

import (
	_ "github.com/gofiber/fiber/v2"
	fiber "github.com/gofiber/fiber/v2"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/internal/model"
)

func Get(c *fiber.Ctx) error {
	{
		var ruangans []model.Ruangan
		// Find all users in database
		result := database.DB.Find(&ruangans)
		// Check for errors during query execution
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return list of users
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Ruangan Berhasil Ditampilkan!",
			"data":    ruangans,
		})
	}
}

func Create(c *fiber.Ctx) error {
	// Parse request body
	{
		var ruangan model.Ruangan
		if err := c.BodyParser(&ruangan); err != nil {
			return err
		}
		// Insert new user into database
		result := database.DB.Create(&ruangan)
		// Check for errors during insertion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Ruangan Berhasil Ditambahkan!",
			"data":    ruangan,
		})
	}
}
func GetRuangan(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("kode_ruangan")
		// Find user by id_user in database
		var ruangan model.Ruangan
		result := database.DB.First(&ruangan, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Ruangan Tidak Ditemukan!",
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
			"data":    ruangan,
		})
	}
}
func Update(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("kode_ruangan")
		// Find user by id_user in database
		var ruangan model.Ruangan
		result := database.DB.First(&ruangan, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Ruangan Tidak Ditemukan",
			})
		}
		// Parse request body
		var newRuangan model.Hari
		if err := c.BodyParser(&newRuangan); err != nil {
			return err
		}
		// Update user in database
		result = database.DB.Model(&ruangan).Updates(newRuangan)
		// Check for errors during update
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Ruangan Berhasil Diperbarui!",
			"data":    ruangan,
		})
	}
}
func Delete(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	{
		id := c.Params("kode_ruangan")
		// Find user by id_user in database
		var ruangan model.Ruangan
		result := database.DB.First(&ruangan, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Ruangan Tidak Ditemukan",
			})
		}
		// Delete user from database
		result = database.DB.Delete(&ruangan)
		// Check for errors during deletion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Ruangan Berhasil Dihapus!",
			"data":    ruangan,
		})
	}
}
