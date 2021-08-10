package repository

import "belajar-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
