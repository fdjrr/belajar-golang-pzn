package golang_database_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "fadjrir'; #"
	password := "salah"

	query := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)

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
