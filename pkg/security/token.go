package security

import (
	"desafio-tecnico-backend/internal/config"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret []byte

func SecretConfig(config *config.Config) error {
	secret = []byte(config.SECRET)

	if len(secret) == 0 {
		return errors.New("token secret is not defined")
	}
	return nil
}

func NewToken(cnpj uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["CNPJ"] = cnpj

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(secret)
}

func GetToken(c *gin.Context) (string, error) {
	const bearer_schema = "Bearer "
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("empty header")
	}

	token := header[len(bearer_schema):]

	err := ValidateToken(token)
	if err != nil {
		return "", errors.New("invalid token")
	}

	return token, nil
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, errors.New("invalid token: " + token)
		}
		return secret, nil
	})
	if err != nil {
		log.Printf("Erro ao validar o token: %s", err)
		return err
	}

	return err

}

func GetCnpjFromToken(token string) (uint64, error) {
	tokenClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, tokenClaims, nil)
	if err != nil {
		return 0, err
	}

	cnpj := tokenClaims["CNPJ"].(uint64)
	return cnpj, nil
}

// func GetPermissions(c *gin.Context) (jwt.MapClaims, error) {
// 	token, err := GetToken(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	permissions, err := ExtractToken(token)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return permissions, nil
// }

// func ExtractToken(tokenString string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenString, keyFunc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	permissions, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, errors.New("error getting permissions")
// 	}

// 	return permissions, nil
// }

// func keyFunc(t *jwt.Token) (interface{}, error) {
// 	_, ok := t.Method.(*jwt.SigningMethodHMAC)
// 	if !ok {
// 		return nil, fmt.Errorf("invalid method: %v,", t.Header["alg"])
// 	}

// 	return secret, nil

// }
