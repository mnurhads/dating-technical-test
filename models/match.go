package models

type Match struct {
	Id     			int    `db:"id" json:"id" gorm:"primaryKey"`
	ProfilId		int    `db:"profil_id" json:"profilId"`
	ProfilTujuan    int    `db:"profil_tujuan" json:"profilTujuan"`
	ProfilMatch     int    `db:"profil_match" json:"profilMatch"`
	Status          string `db:"status" json:"status"`
}

func (b *Match) TableName() string {
	return "matchs"
}