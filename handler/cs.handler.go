package handler

import (
	"tuxedo/services"

	"github.com/gofiber/fiber/v2"
)

func MessageCS(c *fiber.Ctx) error {
	var input struct {
		Message string `json:"message"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	customPrompt := "Tuxedo Labs adalah perusahaan teknologi yang fokus pada pengembangan solusi cloud inovatif untuk berbagai industri. " +
		"Layanan utama kami, Tuxedo Cloud, menawarkan penyimpanan data yang aman, skalabilitas tinggi, dan aksesibilitas global. " +
		"Layanan ini dikembangkan oleh Rafi. Apa yang ingin kamu ketahui lebih lanjut?"

	response, err := services.SendGeminiRequest(customPrompt, input.Message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": response,
	})
}

