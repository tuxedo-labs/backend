package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"tuxedo/database"
	"tuxedo/models/entity"

	"github.com/google/uuid"
)

func GetBlogAll() ([]entity.Blog, error) {
	var blogs []entity.Blog
	err := database.DB.Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func hashFilename(filename string) string {
	ext := filepath.Ext(filename)
	name := uuid.New().String()
	return fmt.Sprintf("%s%s", name, ext)
}

func UploadThumbnail(file *multipart.FileHeader) (string, error) {
	publicDir := "./public/blog/thumbnails"
	filePath := fmt.Sprintf("%s/%s", publicDir, hashFilename(file.Filename))
	relativePath := fmt.Sprintf("/blog/thumbnails/%s", filepath.Base(filePath))

	if err := os.MkdirAll(publicDir, os.ModePerm); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return relativePath, nil
}

func SaveBlog(blog *entity.Blog) error {
	return database.DB.Create(blog).Error
}
