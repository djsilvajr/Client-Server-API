package service

import (
	//"log"

	"my-app/internal/models"
	"my-app/internal/repository"
	"my-app/internal/requests"
	"my-app/internal/response"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var TokenStore = make(map[string]string)
var JwtSecret = []byte("minha-chave-super-secreta-123@2025!bananinha123")

func LoginUser(r requests.LoginUser) (map[string]any, error) {

	var user models.User

	user.Email = r.Email
	user.Password = string(r.Password)

	userSearch, err := repository.ValidateUserLogin(user.Email, user.Password)
	if err != nil {
		return nil, &response.ServiceError{StatusCode: http.StatusUnauthorized, Message: "Usuário ou senha inválidos"}
	}

	// Gerar o token JWT
	token, err := GenerateJWT(user.Email)
	if err != nil {
		return nil, &response.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Erro ao gerar token JWT"}
	}

	TokenStore[user.Email] = token

	//fmt.Printf("%v\n", user)

	retorno := map[string]any{
		"token": token,
		"user": map[string]string{
			"email": r.Email,
			"name":  userSearch.Name,
		},
	}
	return retorno, nil
}

func GenerateJWT(email string) (string, error) {
	// Crie as claims (informações do token)
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // expira em 24h
	}

	// Crie o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assine e gere a string do token
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateServerToken(email, token string) bool {
	savedToken, exists := TokenStore[email]
	return exists && savedToken == token
}
