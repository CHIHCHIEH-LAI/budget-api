package models

type CategoryCreate struct {
	Name   string  `json:"name" binding:"required"`
	Budget float64 `json:"budget" binding:"required"`
	Used   float64 `json:"used" default:"0"`
}

type CategoryUpdate struct {
	Name   string  `json:"name" default:""`
	Budget float64 `json:"budget" default:"0"`
	Used   float64 `json:"used" default:"0"`
}

type Category struct {
	ID     uint    `json:"id" binding:"required"`
	Name   string  `json:"name" binding:"required"`
	Budget float64 `json:"budget" binding:"required"`
	Used   float64 `json:"used" binding:"required"`
}
