package models

// Category represents a budget category
type Category struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name" gorm:"unique;not null"`
	Budget float64 `json:"budget" gorm:"not null;check:budget >= 0"`
	Used   float64 `json:"used" gorm:"not null;default:0;check:used >= 0"`
}
