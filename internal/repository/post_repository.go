package repository

import (
	"fmt"
	"social-sys/internal/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post

	err := r.db.Model(&models.Post{}).First(&post, id).Error

	return &post, err
}

func (r *PostRepository) List(page, pageSize int) ([]models.Post, int, error) {
	var posts []models.Post
	var total int64

	offset := (page - 1) * pageSize
	fmt.Println(pageSize)
	err := r.db.Model(&models.Post{}).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Post{}).
		Limit(pageSize).
		Offset(offset).
		Find(&posts).
		Error

	if err != nil {
		return nil, 0, err
	}
	return posts, int(total), nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Session(&gorm.Session{
		FullSaveAssociations: true,
	}).Save(post).Error
}

func (r *PostRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}
