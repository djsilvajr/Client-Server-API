package repository

import (
	"errors"
	"fmt"
	"my-app/internal/db"
	"my-app/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByID(id uuid.UUID) (*models.User, error) {
	row := db.Conn.QueryRow("SELECT id, name, email, creation_date, update_date FROM users WHERE id = ?", id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreationDate, &user.UpdateDate)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ValidateUserLogin(email string, password string) (*models.User, error) {
	row := db.Conn.QueryRow("SELECT id, name, email, password, creation_date, update_date FROM users WHERE email = ?", email)

	var user models.User
	var hashedPassword string
	err := row.Scan(&user.ID, &user.Name, &user.Email, &hashedPassword, &user.CreationDate, &user.UpdateDate)
	if err != nil {
		return nil, err
	}

	// Verifica se a senha bate com a hash
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// Senha incorreta
		return nil, errors.New("senha inválida")
	}

	fmt.Printf("%v\n", user)
	return &user, nil
}
