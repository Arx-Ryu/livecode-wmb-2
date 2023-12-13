package model

type Discount struct {
	Id				int `gorm:"primaryKey" json:"id"`
	Description		string `json:"description" binding:"required"`
	Pct 			int `json:"pct" binding:"required"`
	CustomerID		[]*Customer `gorm:"many2many:customer_discounts;"`
}

func (Discount) TableName() string {
	return "m_discount" //Ganti nama tabel
}