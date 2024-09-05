package external_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define a struct to match the expected response structure
type GamertagResponse struct {
	Gamertag string `json:"gamertag"`
}

// GetGamertagByXUID fetches the gamertag for a given XUID using the GeyserMC API
func GetGamertagByXUID(xuid string) (string, error) {
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

	body, err := ioutil.ReadAll(res.Body)
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
