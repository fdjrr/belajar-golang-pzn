package repository

import "04-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id int) *entity.Category
}
