package authorcontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/surisalbi/rest-api-go/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var authors []models.Author
	models.DB.Find(&authors)
	return c.Status(fiber.StatusOK).JSON(authors)
}

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author

	if err := models.DB.First(&author, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message" : "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message" : "Data tidak ditemukan",
		})
	}

	return c.JSON(author)
}

func Create(c *fiber.Ctx) error {
	var author models.Author
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := models.DB.Create(&author).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.JSON(author)
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author

	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&author).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "Tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Data berhasil diupdate",
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var author models.Author

	if models.DB.Delete(&author, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message" : "Tidak dapat menghapus data",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Data berhasil dihapus",
	})
}
