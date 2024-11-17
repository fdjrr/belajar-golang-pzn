package repository

import (
	"context"
	"fmt"
	golang_database_mysql "golang-database-mysql"
	"golang-database-mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Name:    "Fadjrir Herlambang",
		Comment: "Belajar Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	ctx := context.Background()

	commentId := 12

	comment, err := commentRepository.FindById(ctx, int32(commentId))

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
