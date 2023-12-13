package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"log"
)

type DiscCRUDUC interface {
	Create(item *model.Discount) (error)
	Read() ([]model.Discount, error)
	Update(oldItem *model.Discount, item model.Discount) (error)
	Delete(item *model.Discount) (error)
}

type discCRUDUC struct {
	repo repository.CustRepo
}

func (u *discCRUDUC) Create(item *model.Discount) (error) {
	return u.repo.CreateDiscount(item)
}

func (u *discCRUDUC) Read() ([]model.Discount, error) {
	items, err := u.repo.FindAllDiscount()
	if err != nil {
		log.Println("Error at FindAll Discount")
		return items, err
	}
	return items, nil
}

func (u *discCRUDUC) Update(oldItem *model.Discount, item model.Discount) (error) {
	return u.repo.UpdateDiscount(oldItem, item)
}

func (u *discCRUDUC) Delete(item *model.Discount) (error) {
	return u.repo.DeleteDiscount(item)
}

func NewDiscCRUDUseCase(repo repository.CustRepo) DiscCRUDUC {
	return &discCRUDUC{
		repo: repo,
	}
}