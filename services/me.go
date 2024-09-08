package services

import (
	"errors"

	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/structs"
)

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
