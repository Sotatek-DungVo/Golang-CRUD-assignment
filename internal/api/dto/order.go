package dto

import "social-sys/internal/models"

type PostListRes struct {
	Posts []models.Post `json:"posts"`
	Total int           `json:"total"`
}
