package usecase

import (
	"server/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) GetById(id int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(id)
	return
}

func (interactor *TodoInteractor) GetAll() (todos []domain.Todo, err error) {
	todos, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) Add(t domain.Todo) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.Store(t)
	return
}

func (interactor *TodoInteractor) Update(t domain.Todo) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.Update(t)
	return
}

func (interactor *TodoInteractor) DeleteById(t domain.Todo) (err error) {
	err = interactor.TodoRepository.DeleteById(t)
	return
}
