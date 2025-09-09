package models

type Member struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Email string `gorm:"type:varchar(50)" json:"email"`
	Telp string `gorm:"type:varchar(20)" json:"telp"`
}