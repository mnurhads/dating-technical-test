package models

type Premium struct {
	Id  	  	int  	 `db:"id" json:"id" gorm:"primaryKey"`
	UserId	  	int  	 `db:"user_id" json:"user_id"`
	Expired     string   `db:"expired" json:"expired"`
}

func (b *Premium) TableName() string {
	return "premium"
}