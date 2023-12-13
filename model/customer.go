package model

import (
	"encoding/json"
	"log"
	"gorm.io/gorm"
)

type Customer struct {
	Id              int `gorm:"primaryKey" json:"id"`
	CustomerName    string `gorm:"size:50;not null" json:"name" binding:"required"`
	MobilePhoneNo   string `json:"phone" binding:"required"`
	IsMember        bool `gorm:"default:false"`
	DiscountID		[]*Discount `gorm:"many2many:customer_discounts;"`
	BillID			[]Bill
	Balance 		int32
	gorm.Model
}

type CustomerDiscounts struct {
	CustomerID		int `gorm:"primaryKey"`
	DiscountID		int `gorm:"primaryKey"`
}

func (Customer) TableName() string {
	return "m_customer" //Ganti nama tabel
}

func (c *Customer) ToString() string {
	customer, err := json.MarshalIndent(c, "", "   ")
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return string(customer)
}