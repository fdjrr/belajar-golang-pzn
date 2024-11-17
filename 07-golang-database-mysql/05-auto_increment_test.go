package golang_database_mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestAutoIncrementTest(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	name := "Ananda Maharani"
	comment := "Belajar Golang Database MySQL!"

	query := "INSERT INTO comments(name, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, query, name, comment)

	if err != nil {
		panic(err)
	}

	commentId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Comment ID:", commentId)
}
