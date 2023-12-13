package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"log"
)

type TableCRUDUC interface {
	Create(item *model.Table) error
	Read() ([]model.Table, error)
	Update(oldItem *model.Table, item model.Table) error
	Delete(item *model.Table) error
}

type tableCRUDUC struct {
	repo repository.BillRepo
}

func (u *tableCRUDUC) Create(item *model.Table) (error) {
	return u.repo.CreateTable(item)
}

func (u *tableCRUDUC) Read() ([]model.Table, error) {
	items, err := u.repo.FindAllTable()
	if err != nil {
		log.Println("Error at FindAll Table")
		return items, err
	}
	return items, nil
}

func (u *tableCRUDUC) Update(oldItem *model.Table, item model.Table) (error) {
	return u.repo.UpdateTable(oldItem, item)
}

func (u *tableCRUDUC) Delete(item *model.Table) (error) {
	return u.repo.DeleteTable(item)
}

func NewTableCRUDUseCase(repo repository.BillRepo) TableCRUDUC {
	return &tableCRUDUC{
		repo: repo,
	}
}