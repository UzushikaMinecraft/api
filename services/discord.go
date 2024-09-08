package services

import (
	"errors"
	"fmt"

	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/structs"
)

func GetDiscordByUUID(uuid string) (*string, error) {
	if uuid == "" {
		return nil, errors.New("uuid is not provided")
	}

	var u structs.DiscordSrvAccounts

	res := db.DiscordSRV.
		Where("uuid == ?", uuid).
		First(&u)

	if res.RowsAffected == 0 {
		return nil, errors.New("no players found")
	}
	if res.Error != nil {
		return nil, errors.New(
			fmt.Sprintf(
				"error occured while processing request: %v",
				res.Error,
			),
		)
	}

	return &u.UUID, nil
}
