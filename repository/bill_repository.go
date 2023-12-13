package repository

import (
	"errors"
	"livecode-wmb-rest-api/model"

	"gorm.io/gorm"
)

type BillRepo interface {
	//CRUD Bill
	Create(newItem *model.Bill) error
	FindById(id int) (model.Bill, error)
	FindAll() ([]model.Bill, error)
	Update(newItem *model.Bill, by model.Bill) error
	Delete(item *model.Bill) error

	//CRUD Table
	CreateTable(newItem *model.Table) error
	FindByIdTable(id int) (model.Table, error)
	FindAllTable() ([]model.Table, error)
	UpdateTable(oldItem *model.Table, by model.Table) error
	UpdateTableAvailability(oldItem *model.Table, by bool) error
	DeleteTable(item *model.Table) error

	//CRUD Trans Type
	CreateTransType(newItem *model.TransType) error
	FindByIdTransType(id int) (model.TransType, error)
	FindAllTransType() ([]model.TransType, error)
	UpdateTransType(oldItem *model.TransType, newItem model.TransType) error
	DeleteTransType(item *model.TransType) error	

	//CRUD Bill Details
	CreateBillDetail(newItem *model.BillDetail) error
	FindByIdBillDetail(id int) (model.BillDetail, error)
	FindAllBillDetail() ([]model.BillDetail, error)
	FindAllBillDetailByBillId(by string) ([]model.BillDetail, error)
	FindSumBill(by string) ([]model.DtoBillDetailMenuPrice, error)
	FindBillDisc(id int) ([]model.Discount, error)
	UpdateBillDetail(newItem *model.BillDetail, by model.BillDetail) error
	DeleteBillDetail(item *model.BillDetail) error
	
	DailySales() ([]model.DtoDailySales, error)
}

type billRepo struct {
	db *gorm.DB
}

func (r *billRepo) Create(newItem *model.Bill) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *billRepo) FindById(id int) (model.Bill, error) {
	var item model.Bill
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

func (r *billRepo) FindAll() ([]model.Bill, error) {
	var items []model.Bill
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

func (r billRepo) Update(oldItem *model.Bill, newItem model.Bill) error {
	result := r.db.Model(oldItem).Updates(newItem).Error
	return result
}

func (r *billRepo) Delete(item *model.Bill) error {
	result := r.db.Delete(item).Error
	return result
}

func (r *billRepo) CreateTable(newItem *model.Table) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *billRepo) FindByIdTable(id int) (model.Table, error) {
	var item model.Table
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

func (r *billRepo) FindAllTable() ([]model.Table, error) {
	var items []model.Table
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

func (r *billRepo) UpdateTable(oldItem *model.Table, by model.Table) error {
	result := r.db.Model(oldItem).Updates(by).Error
	return result
}

func (r *billRepo) UpdateTableAvailability(oldItem *model.Table, by bool) error {
	result := r.db.Model(oldItem).Update("is_available",by).Error
	return result
}

func (r *billRepo) DeleteTable(item *model.Table) error {
	result := r.db.Delete(item).Error
	return result
}

func (r *billRepo) CreateTransType(newItem *model.TransType) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *billRepo) FindByIdTransType(id int) (model.TransType, error) {
	var item model.TransType
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

func (r *billRepo) FindAllTransType() ([]model.TransType, error) {
	var items []model.TransType
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

func (r *billRepo) UpdateTransType(oldItem *model.TransType, newItem model.TransType) error {
	result := r.db.Model(oldItem).Updates(newItem).Error
	return result
}

func (r *billRepo) DeleteTransType(item *model.TransType) error {
	result := r.db.Delete(item).Error
	return result
}

func (r *billRepo) CreateBillDetail(newItem *model.BillDetail) error {
	result := r.db.Create(&newItem).Error
	return result
}

func (r *billRepo) FindByIdBillDetail(id int) (model.BillDetail, error) {
	var item model.BillDetail
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

func (r *billRepo) FindAllBillDetail() ([]model.BillDetail, error) {
	var items []model.BillDetail
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

func (r *billRepo) FindAllBillDetailByBillId(by string) ([]model.BillDetail, error) {
	var items []model.BillDetail
	err := r.db.Where("bill_id=?", by).Find(&items).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return items, err
		}
		return items, err
	} else {
		return items, nil
	}
}

func (r *billRepo) FindSumBill(by string) ([]model.DtoBillDetailMenuPrice, error) {
	var billDetails model.BillDetail
	var result []model.DtoBillDetailMenuPrice
	err := r.db.Model(&billDetails).Select("bill_details.quantity, m_menu_price.price").Joins("LEFT JOIN m_menu_price on bill_details.menu_price_id = m_menu_price.id").Where("bill_details.bill_id=?",by).Find(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return result, err
		}
		return result, err
	} else {
		return result, nil
	}
}

func (r *billRepo) FindBillDisc(id int) ([]model.Discount, error) {
	var result []model.Discount
	err := r.db.Model(&result).Joins("LEFT JOIN customer_discounts on customer_discounts.discount_id = m_discount.id").Where("customer_id=?", id).Find(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return result, err
		}
		return result, err
	} 
	return result, nil	
}

func (r *billRepo) UpdateBillDetail(oldItem *model.BillDetail, newItem model.BillDetail) error {
	result := r.db.Model(oldItem).Updates(newItem).Error
	return result
}

func (r *billRepo) DeleteBillDetail(item *model.BillDetail) error {
	result := r.db.Delete(item).Error
	return result
}

func (r *billRepo) DailySales() ([]model.DtoDailySales, error) {
	var bill model.Bill
	var result []model.DtoDailySales
	err := r.db.Model(&bill).Select("date(t_bill.created_at) as date, SUM(bill_details.quantity * m_menu_price.price) as total").Joins("LEFT JOIN bill_details on bill_details.bill_id = t_bill.id").Joins("LEFT JOIN m_menu_price on bill_details.menu_price_id = m_menu_price.id").Group("date(t_bill.created_at)").Find(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return result, err
		}
		return result, err
	} else {
		return result, nil
	}
}

func NewBillRepository(db *gorm.DB) BillRepo {
	repo := new(billRepo)
	repo.db = db
	return repo
}