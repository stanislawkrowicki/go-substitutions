package discord_service

import "github.com/bwmarrin/discordgo"

const (
	Github = "https://github.com/stanislawkrowicki/go-substitutions"
)

func SourceCode(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: Github,
		},
	})
}
