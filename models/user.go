package models

type User struct {
	Id           int64  `json:"id" gorm:"primary_key"`
	Code         string `json:"code" gorm:"type:varchar(255);unique_index"`
	Name         string `json:"name" gorm:"type:varchar(255)"`
	Username     string `json:"username" gorm:"type:varchar(255);unique_index"`
	Password     string `json:"password" gorm:"type:varchar(255)"`
	Email        string `json:"email" gorm:"type:varchar(255)"`
	Phone        string `json:"phone" gorm:"type:varchar(255)"`
	Age          int    `json:"age" gorm:"type:int(11)"`
	YearBorn     int    `json:"year_born" gorm:"type:int(11)"`
	Token        string `json:"token"`
	PhotoProfile string `json:"photo_profile"`
}
