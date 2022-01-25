package model

type User struct {
	ID            uint       `gorm:"primary_key"`
	Name          string     `gorm:"type:varchar(20);not null;"`
	Age           string     `gorm:"type:varchar(10);not null;"`
	Sex           string     `gorm:"type:varchar(20);not null;"`
}