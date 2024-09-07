package login

import (
	"github.com/ravener/discord-oauth2"
	"github.com/uzushikaminecraft/api/config"
	"golang.org/x/oauth2"
)

var oauthConf *oauth2.Config

func Init() {
	oauthConf = &oauth2.Config{
		Endpoint:     discord.Endpoint,
		Scopes:       []string{discord.ScopeIdentify},
		RedirectURL:  config.Conf.General.CallbackURL,
		ClientID:     config.Conf.Credentials.Discord.ClientID,
		ClientSecret: config.Conf.Credentials.Discord.ClientSecret,
	}
}
