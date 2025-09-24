package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	Password     string    `db:"password,omitempty"` // nunca expor em JSON
	CreationDate time.Time `db:"creation_date" json:"creation_date"`
	UpdateDate   time.Time `db:"update_date" json:"update_date"`
}
