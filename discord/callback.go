package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/structs"
)

// callback endpoint for Discord login
// @Summary callback endpoint for Discord login
// @Description callback endpoint for Discord login
// @Tags login
// @Accept  json
// @Produce  json
// @Param code query string true "Bearer token"
// @Param state query string true "Random state for validating request"
// @Success 200 {array} structs.DiscordUser
// @Failure 500 {object} structs.Error
// @Router /login/callback [get]
func Callback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != config.Conf.Credentials.State {
		return c.Status(400).JSON(structs.Error{
			Error: "state string does not match. are you doing bad thing?",
		})
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(400).JSON(structs.Error{
			Error: "required parameter is not provided",
		})
	}

	cxt := context.Background()

	token, err := oauthConf.Exchange(
		cxt, code,
	)
	if err != nil {
		return c.Status(400).JSON(structs.Error{
			Error: "failed to exchange token",
		})
	}
	if token == nil {
		return c.Status(400).JSON(structs.Error{
			Error: "failed to contact with Discord",
		})
	}

	url := "https://discordapp.com/api/users/@me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(500).JSON(structs.Error{
			Error: "error occured while making request",
		})
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token.AccessToken))

	b, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(500).JSON(structs.Error{
			Error: "error occured while executing request",
		})
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return c.Status(500).JSON(structs.Error{
			Error: fmt.Sprintf("Discord returned status code %v: %v", resp.StatusCode, string(b)),
		})
	}

	var user structs.DiscordUser
	if err := json.Unmarshal(b, &user); err != nil {
		return c.Status(500).JSON(structs.Error{
			Error: fmt.Sprintf("Failed to parse Discord's JSON: %v", err),
		})
	}

	return c.Status(200).JSON(user)
}
