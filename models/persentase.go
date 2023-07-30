package models

type Persentase struct {
	Id        	int  	`db:"id" json:"id" gorm:"primaryKey"`
	UserId 		int  	`db:"user_id" json:"user_id"`
	Like        int     `db:"like" json:"like"`
	Dislike     int     `db:"dislike" json:"dislike"`
	Match       int     `db:"match" json:"match"`
}

func (b *Persentase) TableName() string {
	return "photo"
}