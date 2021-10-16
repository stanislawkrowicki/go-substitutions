package main

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	discordService "go-substitutions/pkg/discord-service"
	"go-substitutions/pkg/env"
	"go-substitutions/pkg/memory"
	"go-substitutions/pkg/tools"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "source",
			Description: "Bot's source code",
		},
		{
			Name:        "changes",
			Description: "Changes for today/tomorrow",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"source":  discordService.SourceCode,
		"changes": discordService.SendSubstitutionsOnCommand,
	}
)

const (
	CantLoadEnv = "Failed to load .env file."
	ErrorText   = "An error occurred. %s"
	ErrorDelay  = 1 * time.Minute
	Delay       = 5 * time.Minute
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	_ = env.LoadEnv()
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

func listen() {
	log.Println("Listening for substitutions...")
	for {
		subst, err := tools.GetSubstitutions("2021-09-29")
		if err != nil {
			log.Printf(ErrorText, err)
			time.Sleep(ErrorDelay)
			continue
		}

		if subst == nil {
			time.Sleep(Delay)
			continue
		}

		exists, err := memory.Exists(*subst)
		if err != nil {
			log.Printf(ErrorText, err)
			time.Sleep(ErrorDelay)
			continue
		}

		if !exists {
			changed, err := memory.Save(*subst)
			if err != nil {
				log.Printf(ErrorText, err)
				time.Sleep(ErrorDelay)
				continue
			}

			discordService.SendSubstitutions(s, subst, changed)
		}

		time.Sleep(Delay)
	}
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

	go listen()

	defer s.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down...")
}
