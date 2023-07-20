package model

import "time"

type Product struct {
	ProductId uint      `json:"product_id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreateAt  time.Time `json:"create_at"`
}
