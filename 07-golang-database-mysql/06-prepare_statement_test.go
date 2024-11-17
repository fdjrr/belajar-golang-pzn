package golang_database_mysql

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(name, comment) VALUES(?, ?)"

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		name := "Fadjrir"
		comment := "Comment " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, name, comment)

		if err != nil {
			panic(err)
		}

		commentId, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID:", commentId)
	}
}
