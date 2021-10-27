package discord_service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go-substitutions/pkg/tools"
	"log"
	"os"
)

const (
	Github                   = "https://github.com/stanislawkrowicki/go-substitutions"
	SubstitutionNewTitle     = "Hey there! On %s there are some changes:"
	SubstitutionChangedTitle = "Hey! Changes on %s have been made:"
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

func SendSubstitutions(s *discordgo.Session, substitutions *tools.Substitutions, changed bool) error {
	channel, err := s.Channel(os.Getenv("DISCORD_CHANNEL_ID"))
	if err != nil {
		return fmt.Errorf("Error while getting messages channel: %v\n", err)
	}

	lastMessage, err := s.ChannelMessage(channel.ID, channel.LastMessageID)
	isLastMessageSubstitution := len(lastMessage.Embeds) > 0 && lastMessage.Author.ID == s.State.User.ID
	if isLastMessageSubstitution && lastMessage.Embeds[0].Title == PrepareTitleString(substitutions.Date, changed) {
		log.Println("These substitutions were already sent to the channel.")
		return nil
	}

	_, err = s.ChannelMessageSendEmbed(channel.ID, PrepareEmbed(changed, substitutions))
	if err != nil {
		return fmt.Errorf("Error while sending substitutions embed: %v\n", err)
	}

	return nil
}
