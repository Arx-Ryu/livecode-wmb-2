package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	Id			int `gorm:"primaryKey" json:"id"`
	MenuName	string `gorm:"size:50;not null" json:"name" binding:"required"`
	MenuPriceID uint `json:"mprice_id" binding:"required"`
	gorm.Model
}

type MenuPrice struct {
	Id				int `gorm:"primaryKey" json:"id"`
	MenuID			[]Menu
	Price 			int `gorm:"not null" json:"price" binding:"required"`
	BillDetailID 	[]BillDetail
	gorm.Model
}

func (Menu) TableName() string {
	return "m_menu" //Ganti nama tabel
}

func (MenuPrice) TableName() string {
	return "m_menu_price" //Ganti nama tabel
}

