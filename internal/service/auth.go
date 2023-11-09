package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tanya-mtv/go-musthave-diploma-tpl.git/internal/models"
)

const (
	signingKey = "kljksj542ds;flks;l;"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo autorisation
}

func NewAuthStorage(repo autorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	user.Salt = RandStr(20)
	user.Password = generatePasswordHash(user.Password, user.Salt)

	return a.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := a.repo.GetUser(username)
	if err != nil {
		return "", err
	}

	inputpass := generatePasswordHash(password, user.Salt)

	if inputpass != user.Password {
		return "", errors.New("unauthorized")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}
