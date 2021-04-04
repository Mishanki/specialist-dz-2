package repositories

import (
	"context"
	"github.com/Mishanki/specialist-dz-2/internal/models"
	"github.com/Mishanki/specialist-dz-2/storage"
)

type User models.RegisterStruct

func (u *User) CreateUserValidation() (string, bool) {
	row, _ := storage.GetDB().Exec(context.Background(), "SELECT id FROM users WHERE username = $1", u.Username)
	if row.RowsAffected() > 0 {
		return "Username is exist", false
	}

	return "", true
}

func (u *User) CreateUser() (uint64, bool) {
	row := storage.GetDB().QueryRow(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2) returning id",
		u.Username, u.Password)

	row.Scan(&u.ID)

	return u.ID, true
}
