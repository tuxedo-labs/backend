package handler

import (
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"os"
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

	geminiAPIKey := os.Getenv("GEMINI_API")
	geminiURL := os.Getenv("GEMINI_URL")

	if geminiAPIKey == "" || geminiURL == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gemini API key or URL is not set",
		})
	}

	client := resty.New()

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": "Data: " + customPrompt,
					},
					{
						"text": "Answer: " + input.Message,
					},
				},
			},
		},
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(geminiURL + geminiAPIKey)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send request to Gemini API",
		})
	}

	var geminiResponse struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := resty.New().JSONUnmarshal(resp.Body(), &geminiResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse Gemini API response",
		})
	}

	if len(geminiResponse.Candidates) > 0 && len(geminiResponse.Candidates[0].Content.Parts) > 0 {
		textResponse := geminiResponse.Candidates[0].Content.Parts[0].Text
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"response": textResponse,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": "No text found in Gemini API response",
	})
}
