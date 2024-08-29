package service

import (
	"social-sys/internal/models"
	"social-sys/internal/repository"
)

type PostServiceInterface interface {
	CreatePost(post *models.Post) error
	GetPost(id uint) (*models.Post, error)
	ListPosts(page, pageSize int) ([]models.Post, int, error)
	UpdatePost(post *models.Post) error
	DeletePost(id uint) error
}

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.repo.Create(post)
}

func (s *PostService) GetPost(id uint) (*models.Post, error) {
	return s.repo.GetByID(id)
}

func (s *PostService) ListPosts(page, pageSize int) ([]models.Post, int, error) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return s.repo.List(page, pageSize)
}

func (s *PostService) UpdatePost(post *models.Post) error {
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(id uint) error {
	return s.repo.Delete(id)
}
