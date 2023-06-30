package models

func (Transaction) TableName() string {
	return "transactions"
}

type Transaction struct {
	Id                  int64                `json:"id" gorm:"primary_key"`
	Code                string               `json:"code" gorm:"type:varchar(255);unique_index"`
	CustomerName        string               `json:"customer_name" gorm:"type:varchar(255)"`
	TransactionDate     string               `json:"transaction_date" gorm:"type:date"`
	TotalPrice          int64                `json:"total_price" gorm:"type:int(11)"`
	CreatedUserID       int64                `json:"created_user_id" gorm:"type:int(11)"`
	CreatedUser         User                 `json:"created_user" gorm:"foreignkey:CreatedUserID"`
	TransactionProducts []TransactionProduct `json:"transaction_products" gorm:"foreignkey:TransactionId"`
}
