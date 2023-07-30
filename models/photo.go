package models

type Photo struct {
	Id        	int  	`db:"id" json:"id" gorm:"primaryKey"`
	UserId 		int  	`db:"user_id" json:"user_id"`
	Image       string	`db:"image" json:"image"`
}

func (b *Photo) TableName() string {
	return "photo"
}
