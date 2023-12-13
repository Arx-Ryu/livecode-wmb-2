package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"log"
)

type MPriceCRUDUC interface {
	Create(item *model.MenuPrice) (error)
	Read() ([]model.MenuPrice, error)
	Update(oldItem *model.MenuPrice, item model.MenuPrice) (error)
	Delete(item *model.MenuPrice) (error)
}

type mPriceCRUDUC struct {
	repo repository.MenuRepo
}

func (u *mPriceCRUDUC) Create(item *model.MenuPrice) (error) {
	return u.repo.CreateMPrice(item)
}

func (u *mPriceCRUDUC) Read() ([]model.MenuPrice, error) {
	items, err := u.repo.FindMPriceAll()
	if err != nil {
		log.Println("Error at FindAll MPrice")
		return items, err
	}
	return items, nil
}

func (u *mPriceCRUDUC) Update(oldItem *model.MenuPrice, item model.MenuPrice) (error) {
	return u.repo.UpdateMPrice(oldItem, item)
}

func (u *mPriceCRUDUC) Delete(item *model.MenuPrice) (error) {
	return u.repo.DeleteMPrice(item)
}

func NewMPriceCRUDUseCase(repo repository.MenuRepo) MPriceCRUDUC {
	return &mPriceCRUDUC{
		repo: repo,
	}
}