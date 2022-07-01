package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"item_code" gorm:"type:varchar(100);column:item_code"`
	Description string `json:"description" gorm:"type:varchar(100);column:description"`
	Quantity    int    `json:"quantity" gorm:"column:quantity"`
	OrderID     uint   `json:"order_id" gorm:"column:order_id"`
}
