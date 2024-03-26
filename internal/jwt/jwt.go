package jwt

import (
	"crypto/rand"
	"encoding/binary"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

func BuildJWTString() (string, error) {

	userID, err := generateUserID()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(configs.TokenExp)),
		},
		UserID: userID,
	})
	tokenString, err := token.SignedString([]byte(configs.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserID(tokenString string) int {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKey), nil
	})

	if err != nil {
		return -1
	}

	if !token.Valid {
		logger.Debug("Token is not valid")
		return -1
	}

	logger.Debug("Token is valid")

	return claims.UserID
}

func generateUserID() (int, error) {

	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		return 0, err
	}

	data := int(binary.BigEndian.Uint32(b))

	return data, nil
}
