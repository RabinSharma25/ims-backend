package dto

import "time"

// UserDto represents the data transfer object for the Users model.
type UserDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ProductDto represents the data transfer object for the Products model.
type ProductDto struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Price         float64  `json:"price"`
	Rating        *float64 `json:"rating,omitempty"`
	StockQuantity int      `json:"stockQuantity"`
}

// SaleDto represents the data transfer object for the Sales model.
type SaleDto struct {
	Id          string    `json:"id"`
	ProductId   string    `json:"productId"`
	Timestamp   time.Time `json:"timestamp"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unitPrice"`
	TotalAmount float64   `json:"totalAmount"`
}

// PurchaseDto represents the data transfer object for the Purchases model.
type PurchaseDto struct {
	Id        string    `json:"id"`
	ProductId string    `json:"productId"`
	Timestamp time.Time `json:"timestamp"`
	Quantity  int       `json:"quantity"`
	UnitCost  float64   `json:"unitCost"`
	TotalCost float64   `json:"totalCost"`
}

// ExpenseDto represents the data transfer object for the Expenses model.
type ExpenseDto struct {
	Id        string    `json:"id"`
	Category  string    `json:"category"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

// SalesSummaryDto represents the data transfer object for the SalesSummary model.
type SalesSummaryDto struct {
	Id               string    `json:"id"`
	TotalValue       float64   `json:"totalValue"`
	ChangePercentage *float64  `json:"changePercentage,omitempty"`
	Date             time.Time `json:"date"`
}

// PurchaseSummaryDto represents the data transfer object for the PurchaseSummary model.
type PurchaseSummaryDto struct {
	Id               string    `json:"id"`
	TotalPurchased   float64   `json:"totalPurchased"`
	ChangePercentage *float64  `json:"changePercentage,omitempty"`
	Date             time.Time `json:"date"`
}

// ExpenseSummaryDto represents the data transfer object for the ExpenseSummary model.
type ExpenseSummaryDto struct {
	Id            string    `json:"id"`
	TotalExpenses float64   `json:"totalExpenses"`
	Date          time.Time `json:"date"`
}

// ExpenseByCategoryDto represents the data transfer object for the ExpenseByCategory model.
type ExpenseByCategoryDto struct {
	Id               string    `json:"id"`
	ExpenseSummaryId string    `json:"expenseSummaryId"`
	Category         string    `json:"category"`
	Amount           string     `json:"amount"`
	Date             time.Time `json:"date"`
}
