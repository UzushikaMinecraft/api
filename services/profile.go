package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/external_api"
	"github.com/uzushikaminecraft/api/structs"
)

func GetProfiles(m map[string]string) (*[]structs.Profile, error) {
	var err error

	// Check if required parameters were provided
	filter, _ := m["filter"]
	sort, sort_ok := m["sort"]
	offset, offset_ok := m["offset"]
	limit, limit_ok := m["limit"]

	// Parse parameters
	// sort
	if !(sort_ok) {
		sort = "asc"
	}
	if !(sort == "desc" || sort == "asc") {
		return nil, errors.New("parameter sort is not valid")
	}

	// order_by
	order_by, order_by_ok := m["order_by"]
	if !(order_by_ok) {
		order_by = "uuid"
	}

	// offset
	var o int
	if offset_ok {
		o, err = strconv.Atoi(offset)
		if err != nil {
			return nil, errors.New("parameter offset is not valid")
		}
	}

	// limit
	var l int
	if limit_ok {
		l, err := strconv.Atoi(limit)
		if err != nil || l < 0 {
			return nil, errors.New("parameter limit is not valid")
		}
		if l > 50 {
			l = 50
		}
	} else {
		l = 50
	}

	if !(order_by == "id" || order_by == "uuid" || order_by == "experience" || order_by == "currency" || order_by == "total_build_blocks" || order_by == "total_destroy_blocks" || order_by == "total_mob_kills" || order_by == "total_play_time") {
		return nil, errors.New("parameter order_by is not valid")
	}

	var profiles *[]structs.Profile
	db.DB.
		Where("uuid LIKE ?", "%"+filter+"%").
		Order(fmt.Sprintf("%v %v", order_by, sort)).
		Offset(o).
		Limit(l).
		Find(&profiles)

	for i, profile := range *profiles {
		var bedrock *structs.Bedrock
		db.DB.
			Where("fuid = ?", profile.UUID).First(&bedrock)

		if bedrock != nil && bedrock.XUID != "" {
			geyserApi := &external_api.GeyserApi{}
			(*profiles)[i].IsBedrock = true
			(*profiles)[i].XUID = bedrock.XUID
			(*profiles)[i].Name, err = geyserApi.GetGamertagByXUID(bedrock.XUID)
			(*profiles)[i].Avatar.Face = "https://uzsk.iamtakagi.net/api/avatar/face/bedrock/" + bedrock.XUID
			(*profiles)[i].Avatar.Head = "https://uzsk.iamtakagi.net/api/avatar/head/bedrock/" + bedrock.XUID
			(*profiles)[i].Avatar.Body = "https://uzsk.iamtakagi.net/api/avatar/body/bedrock/" + bedrock.XUID

			if err != nil {
				return nil, errors.New("error occured while retrieving Bedrock user's skin")
			}
		} else {
			mojangApi := &external_api.MojangApi{}
			(*profiles)[i].IsBedrock = false
			(*profiles)[i].XUID = ""
			(*profiles)[i].Name, err = mojangApi.GetNameByUUID(profile.UUID)
			(*profiles)[i].Avatar.Face = "https://crafatar.com/avatars/" + profile.UUID
			(*profiles)[i].Avatar.Head = "https://crafatar.com/renders/head/" + profile.UUID
			(*profiles)[i].Avatar.Body = "https://crafatar.com/renders/body/" + profile.UUID

			if err != nil {
				return nil, errors.New("error occured while retrieving Java user's skin")
			}
		}
	}

	return profiles, nil
}

func GetProfile(uuid string) (*structs.Profile, error) {
	var err error

	var profile *structs.Profile
	db.DB.Where("uuid = ?", uuid).First(&profile)

	if profile.UUID == "" {
		return nil, errors.New("UUID is not specified")
	}

	var bedrock *structs.Bedrock
	db.DB.Where("fuid = ?", profile.UUID).First(&bedrock)
	if bedrock != nil && bedrock.XUID != "" {
		geyserApi := &external_api.GeyserApi{}
		profile.IsBedrock = true
		profile.XUID = bedrock.XUID
		profile.Name, err = geyserApi.GetGamertagByXUID(bedrock.XUID)
		profile.Avatar.Face = "https://uzsk.iamtakagi.net/api/avatar/face/bedrock/" + bedrock.XUID
		profile.Avatar.Head = "https://uzsk.iamtakagi.net/api/avatar/head/bedrock/" + bedrock.XUID
		profile.Avatar.Body = "https://uzsk.iamtakagi.net/api/avatar/body/bedrock/" + bedrock.XUID

		if err != nil {
			return nil, errors.New("could not find profile")
		}
	} else {
		mojangApi := &external_api.MojangApi{}
		profile.IsBedrock = false
		profile.XUID = ""
		profile.Name, err = mojangApi.GetNameByUUID(profile.UUID)
		profile.Avatar.Face = "https://crafatar.com/avatars/" + profile.UUID
		profile.Avatar.Head = "https://crafatar.com/renders/head/" + profile.UUID
		profile.Avatar.Body = "https://crafatar.com/renders/body/" + profile.UUID

		if err != nil {
			return nil, errors.New("could not find profile")
		}
	}

	return profile, nil
}
