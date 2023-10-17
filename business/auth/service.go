package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type GenerateTokenSpec struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

var jwtKey = []byte("chat-hex-key")

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) GenerateJWT(generateTokenSpec GenerateTokenSpec) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: generateTokenSpec.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *service) ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}