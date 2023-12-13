package repository

import (
	"errors"
	"livecode-wmb-rest-api/model"

	"gorm.io/gorm"
)

type MenuRepo interface {
	//CRUD Menu
	Create(newItem *model.Menu) error
	FindById(id string) (model.Menu, error)
	FindAll() ([]model.Menu, error)
	Update(oldItem *model.Menu, by model.Menu) error
	Delete(item *model.Menu) error

	//CRUD MenuPrice
	CreateMPrice(newItem *model.MenuPrice) error
	FindMPriceById(id string) (model.MenuPrice, error)
	FindMPriceAll() ([]model.MenuPrice, error)
	UpdateMPrice(oldItem *model.MenuPrice, by model.MenuPrice) error
	DeleteMPrice(item *model.MenuPrice) error
}

type menuRepo struct {
	db *gorm.DB
}

func (r *menuRepo) Create(newItem *model.Menu) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *menuRepo) FindById(id string) (model.Menu, error) {
	var item model.Menu
	result := r.db.First(&item, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return item, err
		}
		return item, err
	} else {
		return item, nil
	}
}

func (r *menuRepo) FindAll() ([]model.Menu, error) {
	var items []model.Menu
	result := r.db.Find(&items)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return items, err
		}
		return items, err
	} else {
		return items, nil
	}
}

func (r *menuRepo) Update(oldItem *model.Menu, by model.Menu) error {
	result := r.db.Model(oldItem).Updates(by).Error
	return result
}

func (r *menuRepo) Delete(item *model.Menu) error {
	result := r.db.Delete(item).Error
	return result
}

func (r *menuRepo) CreateMPrice(newItem *model.MenuPrice) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *menuRepo) FindMPriceById(id string) (model.MenuPrice, error) {
	var item model.MenuPrice
	result := r.db.First(&item, "id=?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return item, err
		}
		return item, err
	} else {
		return item, nil
	}
}

func (r *menuRepo) FindMPriceAll() ([]model.MenuPrice, error) {
	var items []model.MenuPrice
	result := r.db.Find(&items)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return items, err
		}
		return items, err
	} else {
		return items, nil
	}
}

func (r *menuRepo) UpdateMPrice(oldItem *model.MenuPrice, by model.MenuPrice) error {
	result := r.db.Model(oldItem).Updates(by).Error
	return result
}

func (r *menuRepo) DeleteMPrice(item *model.MenuPrice) error {
	result := r.db.Delete(item).Error
	return result
}

func NewMenuRepository(db *gorm.DB) MenuRepo {
	repo := new(menuRepo)
	repo.db = db
	return repo
}