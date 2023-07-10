package jamHandler

import (
	_ "github.com/gofiber/fiber/v2"
	fiber "github.com/gofiber/fiber/v2"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/internal/model"
)

func Get(c *fiber.Ctx) error {
	{
		var jams []model.Jam
		// Find all users in database
		result := database.DB.Find(&jams)
		// Check for errors during query execution
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return list of users
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Jam Berhasil Ditampilkan!",
			"data":    jams,
		})
	}
}

func Create(c *fiber.Ctx) error {
	// Parse request body
	{
		var jam model.Jam
		if err := c.BodyParser(&jam); err != nil {
			return err
		}
		// Insert new user into database
		result := database.DB.Create(&jam)
		// Check for errors during insertion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Jam Berhasil Ditambahkan!",
			"data":    jam,
		})
	}
}
func GetJam(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id")
		// Find user by id_user in database
		var jam model.Jam
		result := database.DB.First(&jam, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Jam Tidak Ditemukan!",
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
			"data":    jam,
		})
	}
}
func Update(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id")
		// Find user by id_user in database
		var jam model.Jam
		result := database.DB.First(&jam, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Jam Tidak Ditemukan",
			})
		}
		// Parse request body
		var newJam model.Jam
		if err := c.BodyParser(&newJam); err != nil {
			return err
		}
		// Update user in database
		result = database.DB.Model(&jam).Updates(newJam)
		// Check for errors during update
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Jam Berhasil Diperbarui!",
			"data":    jam,
		})
	}
}
func Delete(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	{
		id := c.Params("id")
		// Find user by id_user in database
		var jam model.Jam
		result := database.DB.First(&jam, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Jam Tidak Ditemukan",
			})
		}
		// Delete user from database
		result = database.DB.Delete(&jam)
		// Check for errors during deletion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Jam Berhasil Dihapus!",
			"data":    jam,
		})
	}
}
