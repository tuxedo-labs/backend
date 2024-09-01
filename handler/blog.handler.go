package handler

import (
	"net/http"
	"time"
	"tuxedo/models/entity"
	"tuxedo/models/request"
	"tuxedo/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetBlog(c *fiber.Ctx) error {
	blogs, err := services.GetBlogAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}
	if len(blogs) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No blogs found",
		})
	}
	response := make([]request.BlogResponse, len(blogs))
	for i, blog := range blogs {
		response[i] = request.BlogResponse{
			ID:          blog.ID.String(),
			Title:       blog.Title,
			Description: blog.Description,
			Thumbnail:   blog.Thumbnail,
			Author: request.AuthorResponse{
				ID:    blog.User.ID,
				Name:  blog.User.Name,
				Email: blog.User.Email,
			},
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
		}
	}

	return c.JSON(fiber.Map{
		"data": response,
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

	response := request.BlogDetailResponse{
		ID:          blog.ID.String(),
		Title:       blog.Title,
		Description: blog.Description,
		Content:     blog.Content,
		Thumbnail:   blog.Thumbnail,
		Author: request.AuthorResponse{
			ID:    blog.User.ID,
			Name:  blog.User.Name,
			Email: blog.User.Email,
		},
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}

func PostBlog(c *fiber.Ctx) error {
	usersInfo := c.Locals("usersInfo")
	claims := usersInfo.(jwt.MapClaims)

	idFloat := claims["id"].(float64)
	userID := uint(idFloat)
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
		Author:      userID,
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
	})
}

func UpdateBlog(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID format",
		})
	}

	usersInfo := c.Locals("usersInfo")
	claims := usersInfo.(jwt.MapClaims)

	idFloat := claims["id"].(float64)
	userID := uint(idFloat)

	var updateData request.UpdateBlogRequest
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
			"error":   err.Error(),
		})
	}

	if file, err := c.FormFile("thumbnail"); err == nil {
		thumbnailPath, err := services.UploadThumbnail(file)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to upload file",
			})
		}
		updateData.Thumbnail = &thumbnailPath
	}

	updateDataMap := map[string]interface{}{}
	if updateData.Title != nil {
		updateDataMap["title"] = *updateData.Title
	}
	if updateData.Description != nil {
		updateDataMap["description"] = *updateData.Description
	}
	if updateData.Content != nil {
		updateDataMap["content"] = *updateData.Content
	}
	if updateData.Thumbnail != nil {
		updateDataMap["thumbnail"] = *updateData.Thumbnail
	}

	updateDataMap["author"] = userID

	blog, err := services.UpdateBlog(id, updateDataMap)
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
