package structs

type LikeDislikeRequest struct {
	UserId   		int  `json:"userId"`
	ProfilId 		int  `json:"profilId"`
	ProfilTujuan	int  `json:"profilTujuan"`
}

type LikeDislikeResponse struct {
	ResponseCode 	int 	`json:"responseCode"`
	ResponseMsg     string  `json:"responseMsg"`
}

type MatchRequest struct {
	Id          int   `json:"Id"`
	UserId      int   `json:"userId"`
	ProfilId    int   `json:"profilId"`
}

