package models

func (Transaction) TableName() string {
	return "transactions"
}

type Transaction struct {
	Id            int64  `json:"id" gorm:"primary_key"`
	Code          string `json:"code" gorm:"type:varchar(255);unique_index"`
	CustomerName  string `json:"customer_name" gorm:"type:varchar(255)"`
	TotalPrice    int64  `json:"total_price" gorm:"type:int(11)"`
	CreatedUserID int64  `json:"created_user_id" gorm:"type:int(11)"`
}
