package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/koesterjannik/starter/logger"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsers() ([]interface{}, error) {
	DescribeUserTable()
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	rows, err := Db.Query(ctx, `SELECT id, email, password FROM "User"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (interface{}, error) {
		var user User
		err := row.Scan(&user.ID, &user.Email, &user.Password)
		return user, err
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

/*
func GetAllUsers() ([]interface{}, error) {
	DescribeUserTable()
	rows, err := Db.Query(context.Background(), `SELECT id, email, password FROM "User"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersSlice []interface{}
	for rows.Next() {
		value, err := rows.Values()

		if err != nil {
			return nil, err
		}
		usersSlice = append(usersSlice, value)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usersSlice, nil
}*/

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
