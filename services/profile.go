package services

import (
	"fmt"
	"strconv"

	"github.com/uzushikaminecraft/api/external_api"
	"github.com/uzushikaminecraft/api/structs"

	"gorm.io/gorm"
)

// Get profiles with query parameters
// @Summary Get profiles
// @Description Get a list of profiles with optional filtering and sorting, etc.
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param filter query string false "Filter criteria" example(550e8400-e29b-41d4-a716-446655440000) default()
// @Param sort query string false "Sort order" example(desc) default(asc)
// @Param offset query int false "Offset for pagination" example(0) default(0)
// @Param limit query int false "Limit for pagination" example(10) default(50)
// @Param order_by query string false "Order by field" example(play_time)
// @Success 200 {array} structs.Profile
// @Router /profiles [get]
func GetProfiles(db *gorm.DB, m map[string]string) *[]structs.Profile {
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
		return nil
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
			return nil
		}
	}

	// limit
	var l int
	if limit_ok {
		l, err := strconv.Atoi(limit)
		if err != nil || l < 0 {
			return nil
		}
		if l > 50 {
			l = 50
		}
	} else {
		l = 50
	}

	if !(order_by == "id" || order_by == "uuid" || order_by == "experience" || order_by == "currency" || order_by == "total_build_blocks" || order_by == "total_destroy_blocks" || order_by == "total_mob_kills" || order_by == "total_play_time") {
		return nil
	}

	var profiles *[]structs.Profile
	db.
		Where("uuid LIKE ?", "%"+filter+"%").
		Order(fmt.Sprintf("%v %v", order_by, sort)).
		Offset(o).
		Limit(l).
		Find(&profiles)

	for i, profile := range *profiles {
		var bedrock *structs.Bedrock
		db.Where("fuid = ?", profile.UUID).First(&bedrock)
		if bedrock != nil && bedrock.XUID != "" {
			geyserApi := &external_api.GeyserApi{}
			(*profiles)[i].IsBedrock = true
			(*profiles)[i].XUID = bedrock.XUID
			(*profiles)[i].Name, err = geyserApi.GetGamertagByXUID(bedrock.XUID)
			(*profiles)[i].Avatar.Face = "https://uzsk.iamtakagi.net/api/avatar/face/bedrock/" + bedrock.XUID
			(*profiles)[i].Avatar.Head = "https://uzsk.iamtakagi.net/api/avatar/head/bedrock/" + bedrock.XUID
			(*profiles)[i].Avatar.Body = "https://uzsk.iamtakagi.net/api/avatar/body/bedrock/" + bedrock.XUID

			if err != nil {
				return nil
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
				return nil
			}
		}
	}

	return profiles
}

// Get profile by UUID
// @Summary Get profile
// @Description Get a profile by UUID
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param uuid path string true "UUID of target profile"
// @Success 200 {object} structs.Profile
// @Router /profiles/{uuid} [get]
func GetProfile(db *gorm.DB, uuid string) *structs.Profile {
	var err error
	
	var profile *structs.Profile
	db.Where("uuid = ?", uuid).First(&profile)

	if profile.UUID == "" {
		return nil
	}

	var bedrock *structs.Bedrock
	db.Where("fuid = ?", profile.UUID).First(&bedrock)
	if bedrock != nil && bedrock.XUID != "" {
		geyserApi := &external_api.GeyserApi{}
		profile.IsBedrock = true
		profile.XUID = bedrock.XUID
		profile.Name, err = geyserApi.GetGamertagByXUID(bedrock.XUID)

		if err != nil {
			return nil
		}
	} else {
		mojangApi := &external_api.MojangApi{}
		profile.IsBedrock = false
		profile.XUID = ""
		profile.Name, err = mojangApi.GetNameByUUID(profile.UUID)

		if err != nil {
			return nil
		}
	}

	return profile
}
