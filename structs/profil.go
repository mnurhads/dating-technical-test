package structs

// import (
// 	"time"
// )

type RequestUpdateProfil struct {
	UserId   	int  			`json:"user_id"` 
	Age      	int   			`json:"age"`
	Birthdate	string   		`json:"birthDate"`
	BirthInfo	string 			`json:"birthInfo"`
	Bio 		string			`json:"bio"`
	Gender   	GenderList 		`json:"gender"`
	Lokasi      string          `json:"lokasi"`
	Image		string		    `json:"imageInfo"`
}

type ResponseUpdateProfil struct {
	ResponseCode 	int 		`json:"responseCode"`
	ResponseMsg     string		`json:"responseMsg"`
	ProfilData		ProfilDataDetail `json:"profilDataDetail"`
}

// request profil
type RequestProfilList struct {
	UserId  int  	`json:"userId"`
	Lokasi  string	`json:"lokasi"`
}

type ResponseProfilLists struct {
	ResponseCode   int 		`json:"responseCode"`
	ResponseMsg    string   `json:"responseMsg"`
	Premium        string   `json:"premium"`
	Profil		   ProfilDataDetails `json:"profil"`
}

type ResponseProfilList struct {
	ResponseCode   int 		`json:"responseCode"`
	ResponseMsg    string   `json:"responseMsg"`
	Profil		   ProfilDataDetails `json:"profil"`
}

type ProfilDataDetails struct {
	Age			int				`json:"age"`
	Birthdate	string			`json:"birthDate"`
	BirthInfo	string 			`json:"birthInfo"`
	Bio 		string			`json:"bio"`
	Gender   	string  		`json:"gender"`
	Image		string		    `json:"image"`
	Lokasi      string          `json:"lokasi"`
}

type ProfilDataDetail struct {
	UserData 	UserDetailData 	`json:"userDetail"`
	Age			int				`json:"age"`
	Birthdate	string			`json:"birthDate"`
	BirthInfo	string 			`json:"birthInfo"`
	Bio 		string			`json:"bio"`
	Gender   	GenderList 		`json:"gender"`
	Image		ImageProfil		`json:"imageInfo"`
}

type UserDetailData struct {
	Username  string	`json:"username"`
	Fullname  string	`json:"fullname"`
	Notelp	  string 	`json:"notelp"`
	Email	  string	`json:"email"`
}

type GenderList struct {
	Kode	string   `json:"kode"`
	Value   int      `json:"value"`
}

type ImageProfil struct {
	Gambar   	string	  `json:"imageProfil"`
}