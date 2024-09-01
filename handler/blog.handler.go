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

func GetBlogByID(c *fiber.Ctx) error {
	id := c.Params("id")
	blog, err := services.GetBlogByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Blog not found",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": blog,
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

// func UpdateBlog(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var updateData map[string]interface{}
// 	if err := c.BodyParser(&updateData); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid data",
// 			"error":   err.Error(),
// 		})
// 	}
//
// 	file, err := c.FormFile("thumbnail")
// 	if err == nil {
// 		thumbnailPath, err := services.UploadThumbnail(file)
// 		if err != nil {
// 			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 				"message": "Failed to upload file",
// 			})
// 		}
// 		updateData["thumbnail"] = thumbnailPath
// 	}
//
// 	blog, err := services.UpdateBlog(id, updateData)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to update blog",
// 			"error":   err.Error(),
// 		})
// 	}
//
// 	return c.JSON(fiber.Map{
// 		"message": "Blog updated successfully",
// 		"data":    blog,
// 	})
// }

func PatchBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
			"error":   err.Error(),
		})
	}

	file, err := c.FormFile("thumbnail")
	if err == nil {
		thumbnailPath, err := services.UploadThumbnail(file)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to upload file",
			})
		}
		updateData["thumbnail"] = thumbnailPath
	}

	blog, err := services.UpdateBlog(id, updateData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update blog",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Blog updated successfully",
		"data":    blog,
	})
}

func DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := services.DeleteBlog(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete blog",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Blog deleted successfully",
	})
}
