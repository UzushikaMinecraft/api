package structs

type DiscordUser struct {
	ID                   string `json:"id"`
	Username             string `json:"username"`
	Avatar               string `json:"avatar"`
	Discriminator        string `json:"discriminator"`
	PublicFlags          int    `json:"public_flags"`
	Flags                int    `json:"flags"`
	Banner               any    `json:"banner"`
	AccentColor          int    `json:"accent_color"`
	GlobalName           string `json:"global_name"`
	AvatarDecorationData any    `json:"avatar_decoration_data"`
	BannerColor          string `json:"banner_color"`
	Clan                 any    `json:"clan"`
	MfaEnabled           bool   `json:"mfa_enabled"`
	Locale               string `json:"locale"`
	PremiumType          int    `json:"premium_type"`
}

type DiscordSRVUser struct {
	Link    int    `gorm:"column:link;primaryKey;autoIncrement" json:"id"`
	Discord string `gorm:"column:discord;not null;type:varchar(32);not null" json:"discord"`
	UUID    string `gorm:"column:uuid;not null;type:varchar(36);not null" json:"uuid"`
}
