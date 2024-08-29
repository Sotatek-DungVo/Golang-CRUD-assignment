package service

import (
	"fmt"
	"social-sys/internal/api/dto"
	"social-sys/internal/models"
	"social-sys/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(createUserinput *dto.CreateUserInput) error
	Login(loginInput *dto.LoginInput) string
}

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(loginInput *dto.LoginInput) string {
	existedUser, err := s.repo.GetByUsername(loginInput.Username)

	if err != nil {

	}

	if existedUser != nil {

	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(loginInput.Password)); err != nil {

	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":  existedUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte("TEMP_SECRET_KEY"))

	if err != nil {

	}

	fmt.Printf("token %v \n", token)
	return token
}

func (s *AuthService) Register(createUserInput *dto.CreateUserInput) error {
	existedUser, err := s.repo.GetByUsername(createUserInput.Username)

	if err != nil {

	}

	if existedUser != nil {

	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(createUserInput.Password), bcrypt.DefaultCost)

	if err != nil {

	}

	userPayload := models.User{
		Name:     createUserInput.Name,
		Password: string(passwordHash),
		Username: createUserInput.Username,
	}

	return s.repo.Create(&userPayload)
}
