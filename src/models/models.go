package models

import "time"

// Users table
// type Users struct {
// 	Id        int64  `gorm:"column:id;primaryKey" json:"id"`
// 	UserName  string `gorm:"column:user_name" json:"userName"`
// 	FirstName string `gorm:"column:first_name" json:"firstName"`
// 	LastName  string `gorm:"column:last_name" json:"lastName"`
// 	Email     string `gorm:"column:email" json:"email"`
// }

// func (Users) TableName() string {
// 	return "public.users"
// }

// Orders table
type Orders struct {
	Id        int64     `gorm:"column:id;primaryKey" json:"id"`
	UserId    int64     `gorm:"column:user_id" json:"userId"`                                 // Foreign Key for Users
	OrderDate time.Time `gorm:"column:order_date;default:current_timestamp" json:"orderDate"` // Timestamp for the order
}

func (Orders) TableName() string {
	return "public.orders"
}

//////////// ****************** ////////////////

// User represents the "users" table.
type Users struct {
	Id    string `gorm:"primaryKey;column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
}

// TableName specifies the table name for the User model.
func (Users) TableName() string {
	return "public.users"
}

// Product represents the "products" table.
type Products struct {
	Id            string   `gorm:"primaryKey;column:id" json:"id"`
	Name          string   `gorm:"column:name" json:"name"`
	Price         float64  `gorm:"column:price" json:"price"`
	Rating        *float64 `gorm:"column:rating" json:"rating,omitempty"`
	StockQuantity int      `gorm:"column:stock_quantity" json:"stockQuantity"`
}

// TableName specifies the table name for the Product model.
func (Products) TableName() string {
	return "public.products"
}

// Sale represents the "sales" table.
type Sales struct {
	Id          string    `gorm:"primaryKey;column:id" json:"id"`
	ProductId   string    `gorm:"column:product_id" json:"productId"`
	Timestamp   time.Time `gorm:"column:timestamp" json:"timestamp"`
	Quantity    int       `gorm:"column:quantity" json:"quantity"`
	UnitPrice   float64   `gorm:"column:unit_price" json:"unitPrice"`
	TotalAmount float64   `gorm:"column:total_amount" json:"totalAmount"`
}

// TableName specifies the table name for the Sale model.
func (Sales) TableName() string {
	return "public.sales"
}

// Purchase represents the "purchases" table.
type Purchases struct {
	Id        string    `gorm:"primaryKey;column:id" json:"id"`
	ProductId string    `gorm:"column:product_id" json:"productId"`
	Timestamp time.Time `gorm:"column:timestamp" json:"timestamp"`
	Quantity  int       `gorm:"column:quantity" json:"quantity"`
	UnitCost  float64   `gorm:"column:unit_cost" json:"unitCost"`
	TotalCost float64   `gorm:"column:total_cost" json:"totalCost"`
}

// TableName specifies the table name for the Purchase model.
func (Purchases) TableName() string {
	return "public.purchases"
}

// Expense represents the "expenses" table.
type Expenses struct {
	Id        string    `gorm:"primaryKey;column:id" json:"id"`
	Category  string    `gorm:"column:category" json:"category"`
	Amount    float64   `gorm:"column:amount" json:"amount"`
	Timestamp time.Time `gorm:"column:timestamp" json:"timestamp"`
}

// TableName specifies the table name for the Expense model.
func (Expenses) TableName() string {
	return "public.expenses"
}

// SalesSummary represents the "sales_summary" table.
type SalesSummary struct {
	Id               string    `gorm:"primaryKey;column:id" json:"id"`
	TotalValue       float64   `gorm:"column:total_value" json:"totalValue"`
	ChangePercentage *float64  `gorm:"column:change_percentage" json:"changePercentage,omitempty"`
	Date             time.Time `gorm:"column:date" json:"date"`
}

// TableName specifies the table name for the SalesSummary model.
func (SalesSummary) TableName() string {
	return "public.sales_summary"
}

// PurchaseSummary represents the "purchase_summary" table.
type PurchaseSummary struct {
	Id               string    `gorm:"primaryKey;column:id" json:"id"`
	TotalPurchased   float64   `gorm:"column:total_purchased" json:"totalPurchased"`
	ChangePercentage *float64  `gorm:"column:change_percentage" json:"changePercentage,omitempty"`
	Date             time.Time `gorm:"column:date" json:"date"`
}

// TableName specifies the table name for the PurchaseSummary model.
func (PurchaseSummary) TableName() string {
	return "public.purchase_summary"
}

// ExpenseSummary represents the "expense_summary" table.
type ExpenseSummary struct {
	Id            string    `gorm:"primaryKey;column:id" json:"id"`
	TotalExpenses float64   `gorm:"column:total_expenses" json:"totalExpenses"`
	Date          time.Time `gorm:"column:date" json:"date"`
}

// TableName specifies the table name for the ExpenseSummary model.
func (ExpenseSummary) TableName() string {
	return "public.expense_summary"
}

// ExpenseByCategory represents the "expense_by_category" table.
type ExpenseByCategory struct {
	Id               string    `gorm:"primaryKey;column:id" json:"id"`
	ExpenseSummaryId string    `gorm:"column:expense_summary_id" json:"expenseSummaryId"`
	Category         string    `gorm:"column:category" json:"category"`
	Amount           int64     `gorm:"column:amount" json:"amount"`
	Date             time.Time `gorm:"column:date" json:"date"`
}

// TableName specifies the table name for the ExpenseByCategory model.
func (ExpenseByCategory) TableName() string {
	return "public.expense_by_category"
}
