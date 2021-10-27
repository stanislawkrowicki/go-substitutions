package discord_service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go-substitutions/pkg/tools"
)

const (
	Color = 16711935
)

func PrepareSubstitutionsString(substitutions *tools.Substitutions) string {
	changes := ""
	for _, change := range substitutions.Changes {
		changes += change + "\n\n"
	}
	return changes
}

func PrepareTitleString(date string, changed bool) string {
	var title string

	if changed {
		title = fmt.Sprintf(SubstitutionChangedTitle, date)
	} else {
		title = fmt.Sprintf(SubstitutionNewTitle, date)
	}

	return title
}

func PrepareEmbed(changed bool, substitutions *tools.Substitutions) *discordgo.MessageEmbed {
	title := PrepareTitleString(substitutions.Date, changed)

	changes := PrepareSubstitutionsString(substitutions)

	embed := discordgo.MessageEmbed{
		URL:         Github,
		Title:       title,
		Description: changes,
		Color:       Color,
	}

	return &embed
}
