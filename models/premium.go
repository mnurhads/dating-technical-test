package models

//import "time"

type Premium struct {
	Id  	  int  	 `db:"id" json:"id" gorm:"primaryKey"`
	UserId	  int  	 `db:"user_id" json:"user_id"`
	ExpiretAt string `db:"expired_at" json:"expiredAt"`
}

func (b *Premium) TableName() string {
	return "premium"
}