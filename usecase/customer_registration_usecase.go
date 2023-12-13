package usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	"livecode-wmb-rest-api/utils"
	"log"
)

type CustomerRegistrationUC interface {
	Register(newCust *model.Customer) error
	GetCustomer() ([]model.Customer, error)
	GetCustomerByName(name string) (model.Customer, error)
	ActivateMember(id int) (model.Customer, error)
}

type customerRegistrationUC struct {
	repo 	repository.CustRepo
}

func (u *customerRegistrationUC) Register(newCust *model.Customer) error {
	return u.repo.Create(newCust)
}

func (u *customerRegistrationUC) GetCustomer() ([]model.Customer, error) {
	cust, err := u.repo.FindAll()
	if utils.IsError(err) {
		log.Println("Failed at FindAll Cust")
		return cust, err
	}
	return cust, nil
}

func (u *customerRegistrationUC) GetCustomerByName(name string) (model.Customer, error) {
	cust, err := u.repo.FindByName(name)
	if utils.IsError(err) {
		log.Println("Failed at FindByName Cust")
		return cust, err
	}
	return cust, nil
}

func (u *customerRegistrationUC) ActivateMember(id int) (model.Customer, error) {
	cust, err := u.repo.FindById(id)
	if utils.IsError(err) {
		log.Println("Failed at FindById Cust")
		return cust, err
	}
	disc, err := u.repo.FindByIdDiscount(1)
	if utils.IsError(err) {
		log.Println("Failed at FindById Disc")
		return cust, err
	}
	newCust := cust
	newCust = model.Customer{
		IsMember: true,
	}
	newMember := model.CustomerDiscounts{
		CustomerID: cust.Id,
		DiscountID: disc.Id,
	}
	err = u.repo.Update(&cust, newCust)
	if utils.IsError(err) {
		log.Println("Failed at Update")
		return newCust, err
	}
	err  = u.repo.CreateCustDisc(&newMember)
	if utils.IsError(err) {
		log.Println("Failed at Create Assoc")
		return newCust, err
	}		
	return u.repo.FindById(cust.Id)
}

func NewCustomerRegistrationUseCase(repo repository.CustRepo) CustomerRegistrationUC {
	return &customerRegistrationUC{
		repo: repo,
	}
}