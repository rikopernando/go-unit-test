package repository

import "golang-unit-test/entity"

type CategoryRepository interface {
	FindById(Id string) *entity.Category
}
