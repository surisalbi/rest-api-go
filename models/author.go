package models

type Author struct {
	Id		int64 `gorm:"primaryKey" json:"id"`
	Name	string `gorm:"type:varchar(50)" json:"name"`
	Email	string `gorm:"type:varchar(50)" json:"email"`
	Address	string `gorm:"type:varchar(255)" json:"address"`
}