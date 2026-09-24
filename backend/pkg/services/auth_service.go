package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/dto"
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/models"
	"github.com/pranav-patidar9354/job-portal-v2/backend/pkg/repositories"
)

type AuthService struct{ Users *repositories.UserRepository }

func (s *AuthService) Register(req dto.RegisterRequest) (*models.User, error) {
	if _, err := s.Users.FindByEmail(req.Email); err == nil {
		return nil, errors.New("email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name: req.Name, Email: req.Email,
		Password: string(hash), Role: req.Role,
	}

	if err := s.Users.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*models.User, string, error) {
	user, err := s.Users.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, "", errors.New("invalid email or password")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, "", errors.New("JWT_SECRET is not configured")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, "", err
	}

	return user, signed, nil
}
