package structs

type PremiumRequest struct {
	UserId     int    `json:"userId"`
	Expired    string `json:"expired"`
}

type PremiumResponse struct {
	ResponseCode int  	`json:"responseCode"`
	ResponseMsg  string	`json:"responseMsg"`
}