package models

type UserLink struct {
	Id  int  `db:"id" json:"id"`
}

func (b *UserLink) TableName() string {
	return "user"
}