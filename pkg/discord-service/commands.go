package discord_service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go-substitutions/pkg/tools"
	"log"
	"os"
	"time"
)

const (
	Github                   = "https://github.com/stanislawkrowicki/go-substitutions"
	SubstitutionNewTitle     = "Hey there! On %s there are some changes:"
	SubstitutionChangedTitle = "Hey! Changes on %s have been made:"
	ErrorTitle               = "Something went wrong."
	ErrorText                = "An error occurred. %s"
	ErrorDelay               = 1 * time.Minute
	Delay                    = 5 * time.Minute
	NoChanges                = "There aren't any changes."
)

func SourceCode(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: Github,
		},
	})
}

func SendSubstitutionsOnCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	substitutions, err := tools.GetSubstitutions(tools.GetRequestDate())
	if err != nil {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Error while fetching changes. %v", err),
			},
		})
		return
	}

	if substitutions == nil {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: NoChanges,
			},
		})
		return
	}

	embed := PrepareEmbed(false, substitutions)

	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

}

func SendSubstitutions(s *discordgo.Session, substitutions *tools.Substitutions, changed bool) {
	_, err := s.ChannelMessageSendEmbed(os.Getenv("DISCORD_CHANNEL_ID"), PrepareEmbed(changed, substitutions))
	if err != nil {
		log.Printf("Error while sending substitutions embed: %v\n", err)
	}
}
