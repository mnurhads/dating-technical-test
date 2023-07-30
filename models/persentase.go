package models

type Persentase struct {
	Id        	int  	`db:"id" json:"id" gorm:"primaryKey"`
	UserId 		int  	`db:"user_id" json:"user_id"`
	Like        int     `db:"like" json:"like"`
	Dislike     int     `db:"dislike" json:"dislike"`
	Matchs      int     `db:"matchs" json:"matchs"`
}

type Like struct {
	Like  int  `db:"like" json:"like"`
}

type Dislike struct {
	Dislike int `db:"dislike" json:"dislike"`
}

type MatchUser struct {
	Matchs  int `db:"matchs" json:"matchs"`
}

func (b *Persentase) TableName() string {
	return "photo"
}