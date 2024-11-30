package models

import "time"

// Users table
type Users struct {
	Id        int64  `gorm:"column:id;primaryKey" json:"id"`
	UserName  string `gorm:"column:user_name" json:"userName"`
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
	Email     string `gorm:"column:email" json:"email"`
}

func (Users) TableName() string {
	return "public.users"
}

// Orders table
type Orders struct {
	Id        int64     `gorm:"column:id;primaryKey" json:"id"`
	UserId    int64     `gorm:"column:user_id" json:"userId"`                                 // Foreign Key for Users
	OrderDate time.Time `gorm:"column:order_date;default:current_timestamp" json:"orderDate"` // Timestamp for the order
}

func (Orders) TableName() string {
	return "public.orders"
}
