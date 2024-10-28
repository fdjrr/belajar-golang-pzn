package repository

import (
	"04-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id int) *entity.Category {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(*entity.Category)

		return category
	}
}
