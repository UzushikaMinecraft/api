package external_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const API_BASEURL = "https://sessionserver.mojang.com"

// MojangApi provides methods to interact with Mojang's API.
type MojangApi struct{}

// UserProfile represents the response structure for username lookup by UUID.
type ProfileResnponse struct {
	Name string `json:"name"`
}

// UUIDResponse represents the response structure for UUID lookup by username.
type UUIDResponse struct {
	ID string `json:"id"`
}

// GetNameByUUID fetches the name associated with the given UUID.
func (api *MojangApi) GetNameByUUID(uuid string) (string, error) {
	endpoint := fmt.Sprintf("%s/session/minecraft/profile/%s", API_BASEURL, strings.ReplaceAll(uuid, "-", ""))
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("error fetching username by UUID: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error! status: %d", resp.StatusCode)
	}

	var profileResponse ProfileResnponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	if err := json.Unmarshal(body, &profileResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return profileResponse.Name, err
}

// GetUUIDByName fetches the UUID associated with the given username.
func (api *MojangApi) GetUUIDByName(name string) (string, error) {
	endpoint := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", name)
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("error fetching UUID by name: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error! status: %d", resp.StatusCode)
	}

	var uuidResponse UUIDResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}
	if err := json.Unmarshal(body, &uuidResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Format UUID (the API might return it without hyphens)
	uuid := uuidResponse.ID
	formattedUUID := fmt.Sprintf("%s-%s-%s-%s-%s", uuid[:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
	return formattedUUID, nil
}
