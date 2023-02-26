package usecase

import (
	"server/domain"
)

type TodoRepository interface {
	FindById(id int) (domain.Todo, error)
	FindAll() ([]domain.Todo, error)
	Store(domain.Todo) (domain.Todo, error)
	Update(domain.Todo) (domain.Todo, error)
	DeleteById(domain.Todo) error
}
