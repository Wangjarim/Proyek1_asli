package belajarHandler

import (
	fiber "github.com/gofiber/fiber/v2"
	"sistem-informasi-klinik/database"
	"sistem-informasi-klinik/internal/model"
	"time"
)

func GetUsers(c *fiber.Ctx) error {
	{
		var users []model.Pasien
		// Find all users in database
		result := database.DB.Preload("JadwalDokter.Dokter").Preload("JadwalDokter.Hari").Preload("JadwalDokter.Jam").Preload("JadwalDokter.Ruangan").Find(&users)
		// Check for errors during query execution
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return list of users
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Data User Berhasil Ditampilkan!",
			"data":    users,
		})
	}
}

func CreateUser(c *fiber.Ctx) error {
	// Parse request body
	var user model.Pasien
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// Set default tanggal reservasi jika nil
	if user.TglReservasi.IsZero() {
		user.TglReservasi = time.Now()
	}

	// Format tanggal lahir sebelum menyimpan ke database
	//user.Tanggallahir = user.Tanggallahir.UTC().Truncate(24 * time.Hour)

	// Insert new user into database
	result := database.DB.Create(&user)
	// Check for errors during insertion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	// Return success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data Berhasil Ditambahkan!",
		"data":    user.Id,
	})
}

func GetUser(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		Id := c.Params("id")
		// Find user by id_user in database
		var user model.Pasien
		result := database.DB.Preload("JadwalDokter.Dokter").Preload("JadwalDokter.Hari").Preload("JadwalDokter.Jam").Preload("JadwalDokter.Ruangan").Where("id = ?", Id).First(&user, Id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User Tidak Ditemukan!",
			})
		}
		// Check for errors during query
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Set default tanggal reservasi jika nil
		if user.TglReservasi.IsZero() {
			user.TglReservasi = time.Now()
		}
		// Return user
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success",
			"data": fiber.Map{
				"id":            user.Id,
				"nama_lengkap":  user.Namalengkap,
				"nik":           user.Nik,
				"jenis_kelamin": user.Jeninkelamin,
				"tempat_lahir":  user.Tempatlahir,
				"tanggal_lahir": user.Tanggallahir,
				"alamat":        user.Alamat,
				"no_hp":         user.Nohp,
				"id_jadwal":     user.IdJadwal,
				"tgl_reservasi": user.TglReservasi,
				"jadwal_dokter": user.JadwalDokter,
			},
		})
	}
}
func UpdateUser(c *fiber.Ctx) error {
	{
		// Get id_user parameter from request url
		id := c.Params("id_user")
		// Find user by id_user in database
		var user model.Pasien
		result := database.DB.First(&user, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User Tidak Ditemukan",
			})
		}
		// Parse request body
		var newUser model.Pasien
		if err := c.BodyParser(&newUser); err != nil {
			return err
		}
		// Update user in database
		result = database.DB.Model(&user).Updates(newUser)
		// Check for errors during update
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User Berhasil Diperbarui!",
			"data":    user,
		})
	}
}
func DeleteUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	{
		id := c.Params("id_user")
		// Find user by id_user in database
		var user model.Pasien
		result := database.DB.First(&user, id)
		// Check if user exists
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User Tidak Ditemukan",
			})
		}
		// Delete user from database
		result = database.DB.Delete(&user)
		// Check for errors during deletion
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
		// Return success message
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User Berhasil Dihapus!",
			"data":    user,
		})
	}
}

func GetJadwalById(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id")
	// Find user by id_user in database
	var user model.Pasien
	result := database.DB.Preload("Pasien.Id").Preload("JadwalDokter.Dokter").Preload("JadwalDokter.Hari").Preload("JadwalDokter.Jam").Preload("JadwalDokter.Ruangan").Where("id = ?", id).First(&user, id)
	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan!",
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
		"data":    user,
	})
}
