package structs

type JWTResponse struct {
	JwtToken string `json:"jwt_token"`
}

type Me struct {
	UserId          string `json:"user_id"`
	SessionExpireAt int64  `json:"session_expire_at"`
}
