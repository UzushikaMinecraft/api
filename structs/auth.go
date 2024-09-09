package structs

type JWTResponse struct {
	Success bool
}

type Me struct {
	Profile         Profile
	SessionExpireAt int64 `json:"session_expire_at"`
}
