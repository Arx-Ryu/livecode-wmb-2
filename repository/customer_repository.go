package repository

import (
	"errors"
	"livecode-wmb-rest-api/model"

	"gorm.io/gorm"
	"context"
	"fmt"
	"log"
	"livecode-wmb-rest-api/service"
)

type CustRepo interface {
	//CRUD Customer
	Create(newItem *model.Customer) error
	FindById(id int) (model.Customer, error)
	FindByName(name string) (model.Customer, error)
	FindAll() ([]model.Customer, error)
	Update(oldItem *model.Customer, by model.Customer) error
	Delete(item *model.Customer) error

	//CRUD Discount
	CreateDiscount(newItem *model.Discount) error
	FindByIdDiscount(id int) (model.Discount, error)
	FindAllDiscount() ([]model.Discount, error)
	UpdateDiscount(oldItem *model.Discount, by model.Discount) error
	DeleteDiscount(item *model.Discount) error

	//C Customer Discount
	CreateCustDisc(newItem *model.CustomerDiscounts) error

	CheckBalance(lopeId int32) (service.ResultMessage, error)
	DoPayment(lopeId int32, amount float32) (bool, error)
}

type custRepo struct {
	db *gorm.DB
	client service.LopeiPaymentClient
}

func (r *custRepo) Create(newItem *model.Customer) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *custRepo) FindById(id int) (model.Customer, error) {
	var item model.Customer
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

func (r *custRepo) FindByName(name string) (model.Customer, error) {
	var item model.Customer
	result := r.db.First(&item, "customer_name=?", name)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return item, err
		}
		return item, err
	} else {
		return item, nil
	}
}

func (r *custRepo) FindAll() ([]model.Customer, error) {
	var items []model.Customer
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

func (r *custRepo) Update(oldItem *model.Customer, by model.Customer) error {
	result := r.db.Model(oldItem).Updates(by).Error
	return result
}

func (r *custRepo) Delete(item *model.Customer) error {
	result := r.db.Delete(item).Error
	return result
}

func (r *custRepo) CreateDiscount(newItem *model.Discount) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *custRepo) FindByIdDiscount(id int) (model.Discount, error) {
	var item model.Discount
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

func (r *custRepo) FindAllDiscount() ([]model.Discount, error) {
	var items []model.Discount
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

func (r *custRepo) UpdateDiscount(oldItem *model.Discount, by model.Discount) error {
	result := r.db.Model(oldItem).Updates(by).Error
	return result
}

func (r *custRepo) DeleteDiscount(item *model.Discount) error {
	result := r.db.Delete(item).Error
	return result
}

func (r *custRepo) CreateCustDisc(newItem *model.CustomerDiscounts) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (c *custRepo) CheckBalance(lopeId int32) (service.ResultMessage, error) {
	balance, err := c.client.CheckBalance(context.Background(), &service.CheckBalanceMessage{LopeiId: lopeId})
	if err != nil {
		log.Fatalln("Error when calling check balance...", err)
	}
	fmt.Println(balance)
	return *balance, err
}

func (c *custRepo) DoPayment(lopeId int32, amount float32) (bool, error) {
	payment, err := c.client.DoPayment(context.Background(), &service.PaymentMessage{
		LopeiId: lopeId,
		Amount:  amount,
	})
	if err != nil {
		log.Fatalln("Error when calling do payment...", err)
		return false, err
	}
	if payment.Result == "FAILED" {
		err = errors.New("insufficient balance")
		return false, err
	}
	return true, err
}

func NewCustomerRepository(db *gorm.DB, client service.LopeiPaymentClient) CustRepo {
	repo := new(custRepo)
	repo.db = db
	repo.client = client
	return repo
}