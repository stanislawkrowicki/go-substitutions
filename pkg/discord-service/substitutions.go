package discord_service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go-substitutions/pkg/tools"
)

const (
	Color = 16711935
)

func PrepareEmbed(changed bool, substitutions *tools.Substitutions) *discordgo.MessageEmbed {
	var title string
	if changed {
		title = fmt.Sprintf(SubstitutionChangedTitle, substitutions.Date)
	} else {
		title = fmt.Sprintf(SubstitutionNewTitle, substitutions.Date)
	}

	changes := ""
	for _, change := range substitutions.Changes {
		changes += change + "\n\n"
	}

	embed := discordgo.MessageEmbed{
		URL:         Github,
		Title:       title,
		Description: changes,
		Color:       Color,
	}

	return &embed
}
