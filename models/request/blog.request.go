package request

import "time"

type UpdateBlogRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Content     *string `json:"content,omitempty"`
	Thumbnail   *string `json:"thumbnail,omitempty"`
}

type AuthorResponse struct {
	ID    uint   `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BlogResponse struct {
	ID          string         `json:"ID"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Thumbnail   string         `json:"thumbnail"`
	Author      AuthorResponse `json:"author"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type BlogDetailResponse struct {
	ID          string         `json:"ID"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Content     string         `json:"content"`
	Thumbnail   string         `json:"thumbnail"`
	Author      AuthorResponse `json:"author"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}
