package main

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	discordService "go-substitutions/pkg/discord-service"
	"go-substitutions/pkg/env"
	"log"
	"os"
	"os/signal"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "source",
			Description: "Bot's source code",
		},
		{
			Name: "changes",
			Description: "Changes for today/tomorrow",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"source": discordService.SourceCode,
		"changes": discordService.SendSubstitutionsOnCommand,
	}
)

const (
	CantLoadEnv = "Failed to load .env file."
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	err := env.LoadEnv()
	if err != nil {
		log.Fatal(CantLoadEnv)
	}
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Failed to open connection. %v", err)
		return
	}

	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, os.Getenv("GUILD_ID"), v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}

	defer s.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down...")
}
