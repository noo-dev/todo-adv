package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaim struct {
	UID int `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, id int) (string, error) {
	// create jwt claims
	expTime := time.Now().Add(time.Hour * 24)
	issTime := time.Now()
	claims := &CustomClaim{
		UID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			Subject:   username,
			IssuedAt:  jwt.NewNumericDate(issTime),
		},
	}

	// add claims to token and define signing method
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign and return jwt string
	token, err := jwt.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(tokenString string, userID int) error {
	parsedToken, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaim{},
		func(jwt *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		})
	if err != nil {
		return err
	}
	if !parsedToken.Valid {
		return errors.New("token is invalid")
	}

	claims, ok := parsedToken.Claims.(*CustomClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.UID != userID {
		return errors.New("permission denied")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}
