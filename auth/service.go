package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int, name string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func (s *jwtService) GenerateToken(userID int, name string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"name":    name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
