package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-substitutions/cmd/memory"
	"go-substitutions/cmd/notifications"
	"go-substitutions/cmd/tools"
	"log"
	"strings"
	"time"
)

const (
	CantLoadEnv              = "Failed to load .env file."
	SubstitutionNewTitle     = "Hey there! On %s there are some changes:"
	SubstitutionChangedTitle = "Hey! Changes on %s have been made:"
	ErrorTitle               = "Something went wrong."
	ErrorText                = "An error occurred. %s"
	ErrorDelay               = 1 * time.Minute
	Delay                    = 5 * time.Minute
)

func main() {
	err := godotenv.Load()
	if err != nil {
		_ = notifications.Show(CantLoadEnv, CantLoadEnv)
		log.Fatal(CantLoadEnv)
	}

	for {
		subst, err := tools.GetSubstitutions(tools.GetRequestDate())
		if err != nil {
			_ = notifications.Show(ErrorTitle, fmt.Sprintf(ErrorText, err))
			time.Sleep(ErrorDelay)
			continue
		}

		if subst == nil {
			time.Sleep(Delay)
			continue
		}

		exists, err := memory.Exists(*subst)
		if err != nil {
			_ = notifications.Show(ErrorTitle, fmt.Sprintf(ErrorText, err))
			time.Sleep(ErrorDelay)
			continue
		}

		if !exists {
			changed, err := memory.Save(*subst)
			if err != nil {
				_ = notifications.Show(ErrorTitle, fmt.Sprintf(ErrorText, err))
				time.Sleep(ErrorDelay)
				continue
			}

			if changed {
				title := fmt.Sprintf(SubstitutionChangedTitle, subst.Date)
				_ = notifications.Show(title, strings.Join(subst.Changes, "\n"))
			} else {
				title := fmt.Sprintf(SubstitutionNewTitle, subst.Date)
				_ = notifications.Show(title, strings.Join(subst.Changes, "\n"))
			}
		}

		time.Sleep(Delay)
	}

}
