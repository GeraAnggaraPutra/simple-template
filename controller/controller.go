package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsersController(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []User

		query := "SELECT id, email FROM users"

		rows, err := db.Query(query)
		if err != nil {
			return err
		}

		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Email)
			if err != nil {
				return err
			}
			users = append(users, user)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": users,
		})
	}
}
