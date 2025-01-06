package models

type CategoryCreate struct {
	Name    string  `json:"name" binding:"required"`
	Budget  float64 `json:"budget" binding:"required,gte=0"`
	Expense float64 `json:"expense" default:"0" binding:"gte=0"`
}

type CategoryBudgetUpdate struct {
	Budget float64 `json:"budget" binding:"required,gte=0"`
}

type CategoryExpenseUpdate struct {
	Expense float64 `json:"expense" binding:"required,gte=0"`
}

type Category struct {
	ID      uint    `json:"id" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Budget  float64 `json:"budget" binding:"required,gte=0"`
	Expense float64 `json:"expense" binding:"required,gte=0"`
}
