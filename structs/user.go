package structs

type RegisterRequest struct {
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Notelp      string `json:"notelp"`
	Password    string `json:"password"`
	Confirm     string `json:"confirm"`
}

type RegisterResponse struct {
	ResponseCode  int `json:"responseCode"`
	ResponseMsg   string `json:"responseMsg"`
}

type LoginRequest struct {
	Username  string  `json:"username" validate:"required,max=100"`
	Password  string  `json:"password" validate:"required,max=100"`
}

type LoginResponse struct {
	ResponseCode int 			`json:"responseCode"`
	ResponseMsg  string 		`json:"responseMsg"`
	Data         Data			`json:"data"`
}

type Data struct {
	TokenData    ResponseToken  `json:"accessTokenData"`
	UserData     UserDetail 	`json:"userData"`
}

type ResponseToken struct {
	AccesToken		string 	`json:"acces_token"`
	TokenType		string	`json:"token_type"`
	ExpiresIn		string	`json:"expires_in"`
}

type UserDetail struct {
	Username 	string `json:"username"`
	Fullname	string `json:"fullname"`
	Email       string `json:"email"`
	Notelp      string `json:"notelp"`
	Status      string `json:"status"`
}

type ErrorResponse struct {
	ResponseCode  int `json:"responseCode"`
	ResponseMsg   string `json:"responseMsg"`
}