package models

type Product struct {
	Id int64 `json:"id" gorm:"primary_key"`
	// add code field with unique index
	Code        string `json:"code" gorm:"type:varchar(255);unique_index"`
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Status      bool   `json:"status" gorm:"type:tinyint(1)"`
	Age 	   int    `json:"age" gorm:"type:int(11)"`
	YearBorn   int    `json:"year_born" gorm:"type:int(11)"`
	// add created user id column with reference to user table
	CreatedUserId int64 `json:"created_user_id" gorm:"type:int(11)"`
	// add updated user id column with reference to user table
	CreatedByUser User
}



type Model struct {
    ID             int
    Name           string
    CreatedUserID int

    // Embed the User struct to establish the relation
    
}