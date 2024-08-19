package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo/internal/models"
	"todo/internal/repository"
)

const (
	salt      = "b52636835ce149418d26ad4f45f11023"
	signature = "d626828560b948d095d8fb9057dd6614"
)

type AuthService struct {
	repository repository.Auth
}

type UserClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repository repository.Auth) *AuthService {
	return &AuthService{repository: repository}
}

func (service *AuthService) CreateUser(user models.UserSignUpInput) (int, error) {
	user.Password = hashPassword(user.Password)

	return service.repository.CreateUser(user)
}

func (service *AuthService) GenerateToken(user models.UserSignInInput) (string, error) {
	existingUser, err := service.repository.GetUser(user.Email)
	if err != nil {
		return "", err
	}

	if hashPassword(user.Password) != existingUser.Password {
		return "", errors.New("password is incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
		UserId: existingUser.Id,
	})

	return token.SignedString([]byte(signature))
}

func (service *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signature), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}

	return claims.UserId, nil
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
