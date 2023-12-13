package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"log"
)

type MenuCRUDUC interface {
	Create(item *model.Menu) (error)
	Read() ([]model.Menu, error)
	Update(oldItem *model.Menu, item model.Menu) (error)
	Delete(item *model.Menu) (error)
}

type menuCRUDUC struct {
	repo repository.MenuRepo
}

func (u *menuCRUDUC) Create(item *model.Menu) (error) {
	return u.repo.Create(item)
}

func (u *menuCRUDUC) Read() ([]model.Menu, error) {
	items, err := u.repo.FindAll()
	if err != nil {
		log.Println("Error at FindAll Table")
		return items, err
	}
	return items, nil
}

func (u *menuCRUDUC) Update(oldItem *model.Menu, item model.Menu) (error) {
	return u.repo.Update(oldItem, item)
}

func (u *menuCRUDUC) Delete(item *model.Menu) (error) {
	return u.repo.Delete(item)
}

func NewMenuCRUDUseCase(repo repository.MenuRepo) MenuCRUDUC {
	return &menuCRUDUC{
		repo: repo,
	}
}