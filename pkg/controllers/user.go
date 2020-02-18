package controllers

import (
	"database/sql"

	"github.com/pojntfx/go-support/pkg/models"
)

type UserController struct {
	DB *sql.DB
}

func (c *UserController) FindById(id int) (error, models.UserModel) {
	user := models.UserModel{}
	c.DB.QueryRow("SELECT id, email, firstname, secondname, username, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.FirstName, &user.SecondName, &user.UserName, &user.Password)

	return nil, user
}

func (c *UserController) FindAll() (error, []models.UserModel) {
	users := []models.UserModel{}

	rows, err := c.DB.Query("SELECT id, email, firstname, secondname, username, password FROM users")
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	for rows.Next() {
		user := models.UserModel{}

		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.SecondName, &user.UserName, &user.Password); err != nil {
			return err, nil
		}

		users = append(users, user)
	}

	err = rows.Err()

	return err, users
}
