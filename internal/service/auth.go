package service

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"my-app/internal/models"
	"my-app/internal/repository"
	"my-app/internal/requests"
	"my-app/internal/response"

	"github.com/golang-jwt/jwt/v5"
)

// tokenStore armazena tokens por email. Protegido por mutex para concorrência.
var (
	tokenStore   = make(map[string]string)
	tokenStoreMu sync.RWMutex

	// JwtSecret é lido uma vez no init; se não definido, usa um fallback (não recomendado para produção).
	JwtSecret = func() []byte {
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			// aviso em runtime é responsabilidade de quem inicializa a aplicação, mas aqui garantimos um fallback
			// Em produção, prefira falhar rápido se a variável não estiver setada.
			return []byte("change-me-in-production")
		}
		return []byte(secret)
	}()
	// TTL padrão do token
	tokenTTL = 24 * time.Hour
)

// LoginUser realiza validação do usuário e gera um JWT se autenticado.
// Retorna um map[string]any com token e dados do usuário, ou um *response.ServiceError para erros tratados.
func LoginUser(r requests.LoginUser) (map[string]any, error) {
	// Monta o modelo de usuário para validação (a conversão/checagem de senha fica no repository)
	user := models.User{
		Email:    r.Email,
		Password: string(r.Password),
	}

	userSearch, err := repository.ValidateUserLogin(user.Email, user.Password)
	if err != nil {
		// Retorna erro de autenticação de forma tratada
		return nil, &response.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Usuário ou senha inválidos",
		}
	}

	// Gera token JWT
	token, err := GenerateJWT(user.Email)
	if err != nil {
		return nil, &response.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Erro ao gerar token JWT",
		}
	}

	// Armazena token de forma concorrente segura
	tokenStoreMu.Lock()
	tokenStore[user.Email] = token
	tokenStoreMu.Unlock()

	retorno := map[string]any{
		"token": token,
		"user": map[string]string{
			"email": r.Email,
			"name":  userSearch.Name,
		},
	}
	return retorno, nil
}

// GenerateJWT cria um JWT assinado com HS256 e expiração definida por tokenTTL.
func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(tokenTTL).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

// ValidateServerToken verifica se o token fornecido coincide com o token armazenado para um email.
func ValidateServerToken(email, token string) bool {
	tokenStoreMu.RLock()
	defer tokenStoreMu.RUnlock()
	saved, ok := tokenStore[email]
	return ok && saved == token
}
