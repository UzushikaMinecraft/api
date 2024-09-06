package external_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeyserApi struct{}

type GamertagResponse struct {
	Gamertag string `json:"gamertag"`
}

type XUIDResponse struct {
	XUID string `json:"xuid"`
}

type SkinResponse struct {
	Hash       string `json:"hash"`
	IsSteve    bool   `json:"is_steve"`
	LastUpdate int64  `json:"last_update"`
	Signature  string `json:"signature"`
	TextureID  string `json:"texture_id"`
	Value      string `json:"value"`
}

// GetGamertagByXUID fetches the gamertag for a given XUID using the GeyserMC API
func (api *GeyserApi) GetGamertagByXUID(xuid string) (string, error) {
	url := fmt.Sprintf("https://api.geysermc.org/v2/xbox/gamertag/%s", xuid)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response into the GamertagResponse struct
	var response GamertagResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.Gamertag, nil
}

func (api *GeyserApi) GetXUIDbyGamertag(xuid string) (string, error) {
	url := fmt.Sprintf("https://api.geysermc.org/v2/xbox/xuid/%s", xuid)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response into the GamertagResponse struct
	var response XUIDResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.XUID, nil
}

// https://api.geysermc.org/v2/skin/{xuid}
func (api *GeyserApi) GetSkinByXUID(xuid string) (*SkinResponse, error) {
	url := fmt.Sprintf("https://api.geysermc.org/v2/skin/%s", xuid)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into the SkinResponse struct
	var response *SkinResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
