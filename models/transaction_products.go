package models

func (TransactionProduct) TableName() string {
	return "transaction_products"
}

type TransactionProduct struct {
	TransactionId int64 `json:"transaction_id" gorm:"type:int(11)" sql:"unique_index:transaction_product"`
	ProductId     int64 `json:"product_id" gorm:"type:int(11)" sql:"unique_index:transaction_product"`

	Quantity     int64  `json:"quantity" gorm:"type:int(11)"`
	ProductName  string `json:"product_name" gorm:"type:varchar(255)"`
	ProductPrice int64  `json:"product_price" gorm:"type:int(11)"`
	SubTotal     int64  `json:"sub_total" gorm:"type:int(11)"`
}
