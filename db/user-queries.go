package db

import (
	"context"

	"github.com/koesterjannik/starter/logger"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsers() ([]User, error) {
	DescribeUserTable()
	rows, err := Db.Query(context.Background(), `SELECT id, email, password FROM "User"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func DescribeUserTable() (string, error) {
	rows, err := Db.Query(context.Background(), "SELECT column_name, data_type FROM information_schema.columns WHERE table_name = 'User'")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var description string
	for rows.Next() {
		var columnName, dataType string
		err = rows.Scan(&columnName, &dataType)
		if err != nil {
			return "", err
		}
		description += columnName + " " + dataType + "\n"
	}

	if err = rows.Err(); err != nil {
		return "", err
	}
	logger.Logger.Info("User table description: " + description)

	return description, nil
}
