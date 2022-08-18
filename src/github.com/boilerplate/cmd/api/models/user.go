package models

import (
	"context"

	"github.com/boilerplate/pkg/application"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

func (u *User) Create(ctx context.Context, app *application.Application) error {
	stmt := `
		INSERT INTO users (
			username
		)
		VALUES ($1)
		RETURNING id
	`

	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		u.UserName,
	).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetByID(ctx context.Context, app *application.Application) error {
	stmt := `
		SELECT *
		FROM users
		WHERE id = $1
	`

	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		u.ID,
	).Scan(
		&u.ID,
		&u.UserName,
	)

	if err != nil {
		return err
	}

	return nil
}

// https://go.dev/doc/database/querying
func GetAllUser(ctx context.Context, app *application.Application) ([]User, error) {
	rows, err := app.DB.Client.Query("select * from users")

	if err != nil {
		return nil, err
	}

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.UserName); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}
