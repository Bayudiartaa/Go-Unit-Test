package repository

import "Go-unit-test/entity"

type CategoryRepository interface{
	FindById(id string) *entity.Category
}