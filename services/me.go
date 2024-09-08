package services

import (
	"errors"

	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/db"
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

func UpdateBiography(biography string, token string) (*structs.Me, error) {
	if token == "" {
		return nil, errors.New("token is not provided")
	}

	claims, err := auth.Validate(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	var p structs.Profile
	p.UUID = claims["user_id"].(string)

	res := db.Core.First(&p)
	if res.RowsAffected == 0 {
		return nil, errors.New("no player found")
	}

	p.Biography = biography
	db.Core.Save(&p)

	return &structs.Me{
		UserId:          claims["user_id"].(string),
		SessionExpireAt: claims["exp"].(int64),
	}, nil
}
