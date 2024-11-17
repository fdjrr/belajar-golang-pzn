package golang_database_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSqlParam(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	name := "Fadjrir Herlambang"
	balance := 10000
	birthDate := "2004-12-05"
	rating := 95.0

	query := "INSERT INTO customers(name, balance, birth_date, rating) VALUES(?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, name, balance, birthDate, rating)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySqlParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "fadjrir'; #"
	password := "salah"

	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"

	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success login", username)
	} else {
		fmt.Println("Gagal login")
	}
}
