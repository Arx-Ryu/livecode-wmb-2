package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"log"
)

type TransTypeCRUDUC interface {
	Create(item *model.TransType) (error)
	Read() ([]model.TransType, error)
	Update(oldItem *model.TransType, item model.TransType) (error)
	Delete(item *model.TransType) (error)
}

type transTypeCRUDUC struct {
	repo repository.BillRepo
}

func (u *transTypeCRUDUC) Create(item *model.TransType) (error) {
	return u.repo.CreateTransType(item)
}

func (u *transTypeCRUDUC) Read() ([]model.TransType, error) {
	items, err := u.repo.FindAllTransType()
	if err != nil {
		log.Println("Error at FindAll Table")
		return items, err
	}
	return items, nil
}

func (u *transTypeCRUDUC) Update(oldItem *model.TransType, item model.TransType) (error) {
	return u.repo.UpdateTransType(oldItem, item)
}

func (u *transTypeCRUDUC) Delete(item *model.TransType) (error) {
	return u.repo.DeleteTransType(item)
}

func NewTransTypeCRUDUseCase(repo repository.BillRepo) TransTypeCRUDUC {
	return &transTypeCRUDUC{
		repo: repo,
	}
}