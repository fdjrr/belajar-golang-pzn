package golang_database_mysql

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestDbTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// begin
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	query := "INSERT INTO comments(name, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		name := "Rani"
		comment := "Comment " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, name, comment)

		if err != nil {
			panic(err)
		}

		commentId, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID:", commentId)
	}

	// commit
	err = tx.Commit()

	if err != nil {
		panic(err)
	}
}
