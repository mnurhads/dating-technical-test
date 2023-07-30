package models

import "time"

type Profil struct {
	Id        	int  	`db:"id" json:"id" gorm:"primaryKey"`
	UserId 		int  	`db:"user_id" json:"user_id"`
	Gender      int		`db:"gender" json:"gender"`
	Age         int     `db:"age" json:"age"`
	Birthdate   time.Time `db:"birthdate" json:"birthdate"`
	BirthInfo   string  `db:"birth_info" json:"birthInfo"`
	Bio   		string  `db:"bio" json:"bio"`
}

func (b *Profil) TableName() string {
	return "photo"
}