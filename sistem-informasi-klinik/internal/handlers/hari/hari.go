package hariHandler

import (
	_ "github.com/gofiber/fiber/v2"
	fiber "github.com/gofiber/fiber/v2"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/internal/model"
)

func Get(c *fiber.Ctx) error {
	{
		var days []model.Hari
		// Find all users in database
		result := database.DB.Find(&days)
		// Check for errors during query execution
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return list of users
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hari Berhasil Ditampilkan!",
			"data":    days,
		})
	}
}

func Create(c *fiber.Ctx) error {
	// Parse request body
	{
		var day model.Hari
		if err := c.BodyParser(&day); err != nil {
			return err
		}
		// Insert new user into database
		result := database.DB.Create(&day)
		// Check for errors during insertion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Hari Berhasil Ditambahkan!",
			"data":    day,
		})
	}
}
func GetHari(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id")
		// Find user by id_user in database
		var day model.Hari
		result := database.DB.First(&day, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Hari Tidak Ditemukan!",
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
			"data":    day,
		})
	}
}
func Update(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id")
		// Find user by id_user in database
		var day model.Hari
		result := database.DB.First(&day, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Hari Tidak Ditemukan",
			})
		}
		// Parse request body
		var newDay model.Hari
		if err := c.BodyParser(&newDay); err != nil {
			return err
		}
		// Update user in database
		result = database.DB.Model(&day).Updates(newDay)
		// Check for errors during update
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hari Berhasil Diperbarui!",
			"data":    day,
		})
	}
}
func Delete(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	{
		id := c.Params("id")
		// Find user by id_user in database
		var day model.Hari
		result := database.DB.First(&day, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Hari Tidak Ditemukan",
			})
		}
		// Delete user from database
		result = database.DB.Delete(&day)
		// Check for errors during deletion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hari Berhasil Dihapus!",
			"data":    day,
		})
	}
}
