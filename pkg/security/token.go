package security

import (
	"desafio-tecnico-backend/internal/config"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret []byte

// Configura o segredo de acordo com as configurações do sistema
func SecretConfig(config *config.Config) error {
	secret = []byte(config.SECRET)

	if len(secret) == 0 {
		return errors.New("token secret is not defined")
	}
	return nil
}

// Cria um novo token de autenticação JWT e retorna como string
func NewToken(cnpj uint64) (string, error) {

	cnpjStr := strconv.FormatUint(cnpj, 10)

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["CNPJ"] = cnpjStr

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(secret)
}

// Pega o token de autenticação em forma de string
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

// Valida um token de autenticação JWT em formato string
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

// Retorna o CNPJ contido nas Claims do token JWT em formato de string
func GetUserCNPJ(c *gin.Context) (string, error) {

	tokenString, _ := GetToken(c)

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("error parsing the token")
		return "", err
	}

	if !token.Valid {
		log.Println("Invalid JWT token")
	}

	cnpj := claims["CNPJ"].(string)

	return string(cnpj), nil

}
