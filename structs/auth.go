package structs

type JWTResponse struct {
	Success bool
}

type Me struct {
	UserId          string `json:"user_id"`
	SessionExpireAt int64  `json:"session_expire_at"`
}
