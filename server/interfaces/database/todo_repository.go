package database

import (
	"server/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(id int) (todo domain.Todo, err error) {
	if err = repo.Find(&todo, id).Error; err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindAll() (todos []domain.Todo, err error) {
	if err = repo.Find(&todos).Error; err != nil {
		return
	}
	return
}

func (repo *TodoRepository) Store(t domain.Todo) (todo domain.Todo, err error) {
	if err = repo.Create(&t).Error; err != nil {
		return
	}
	todo = t
	return
}

func (repo *TodoRepository) Update(t domain.Todo) (todo domain.Todo, err error) {
	if err = repo.Save(&t).Error; err != nil {
		return
	}
	todo = t
	return
}

func (repo *TodoRepository) DeleteById(t domain.Todo) (err error) {
	if err = repo.Delete(&t).Error; err != nil {
		return
	}
	return
}
