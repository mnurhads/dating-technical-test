package models

type User struct {
	Id        	    int     `db:"id" json:"id" gorm:"primaryKey"`
	Username 		string  `db:"username" json:"username"`
	Fullname 		string  `db:"fullname" json:"fullname"`
	Email 			string  `db:"email" json:"email"`
	Notelp  		string `db:"notelp" json:"notelp"`
	Status  		string `db:"status" json:"status"`
	Token  			string `db:"token" json:"token"`
	Password        string `db:"password" json:"password"`
}

func (b *User) TableName() string {
 	return "user"
}

