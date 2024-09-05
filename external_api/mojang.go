package external_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MojangApi struct{}

type NameResponse struct {
	Name string `json:"name"`
}

type UUIDResponse struct {
	UUID string `json:"id"`
}

// Resolve Java players' name by UUID
func (api *MojangApi) GetNameByUUID(uuid string) (string, error) {
	endpoint := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", strings.ReplaceAll(uuid, "-", ""))

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var profileResponse NameResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &profileResponse); err != nil {
		return "", err
	}

	return profileResponse.Name, err
}

// Resolve Java players' UUID by name
func (api *MojangApi) GetUUIDByName(name string) (string, error) {
	endpoint := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", name)
	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var uuidResponse UUIDResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &uuidResponse); err != nil {
		return "", err
	}

	// Format UUID (the API might return it without hyphens)
	uuid := uuidResponse.UUID
	formattedUUID := fmt.Sprintf("%s-%s-%s-%s-%s", uuid[:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
	return formattedUUID, nil
}
