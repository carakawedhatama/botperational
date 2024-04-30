package discord

// reference : https://birdie0.github.io/discord-webhooks-guide/discord_webhook.html
type Discord struct {
	AvatarUrl string          `json:"avatar_url,omitempty"`
	Content   string          `json:"content"`
	Embeds    []DiscordEmbeds `json:"embeds,omitempty"`
}

type DiscordEmbeds struct {
	Description string         `json:"description,omitempty"`
	Image       DiscordImages  `json:"image,omitempty"`
	Footer      DiscordFooters `json:"footer,omitempty"`
}

type DiscordImages struct {
	Url string `json:"url,omitempty"`
}

type DiscordFooters struct {
	Text string `json:"text,omitempty"`
}
