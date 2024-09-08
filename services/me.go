package services

import (
	"errors"

	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/structs"
)

// retrieve information of authenticated user
// @Summary retrieve information of authenticated user
// @Description retrieve information of authenticated user
// @Tags login
// @Accept json
// @Param X-Auth-Token header string true "JSON Web Token"
// @Success 200 {object} structs.Me
// @Failure 400 {object} structs.Error
// @Router /me [get]
func GetMe(token string) (*structs.Me, error) {
	if token == "" {
		return nil, errors.New("token is not provided")
	}

	claims, err := auth.Validate(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return &structs.Me{
		UserId:          claims["user_id"].(string),
		SessionExpireAt: claims["exp"].(int64),
	}, nil
}
