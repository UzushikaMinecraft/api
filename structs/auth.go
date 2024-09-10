package structs

import "github.com/golang-jwt/jwt"

type JWTResponse struct {
	Success bool
}

type JWTCallback struct {
	Claims jwt.MapClaims
	AccessToken string
}

type Me struct {
	Profile         Profile
	SessionExpireAt int64 `json:"session_expire_at"`
}
