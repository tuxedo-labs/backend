package handler

import (
	"net/http"
	"time"
	"tuxedo/models/entity"
	"tuxedo/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetBlog(c *fiber.Ctx) error {
	data, err := services.GetBlogAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	if len(data) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No blogs found",
		})
	}
	return c.JSON(fiber.Map{
		"data": data,
	})
}

func PostBlog(c *fiber.Ctx) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	content := c.FormValue("content")
	file, err := c.FormFile("thumbnail")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid file",
		})
	}

	thumbnailPath, err := services.UploadThumbnail(file)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload file",
		})
	}

	blog := &entity.Blog{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Content:     content,
		Thumbnail:   thumbnailPath,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := services.SaveBlog(blog); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save blog",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Blog created successfully",
		"data":    blog,
	})
}
