module go-substitutions

// +heroku goVersion go1.17
go 1.17

// +heroku install ./cmd/discord-service

require (
	github.com/antchfx/htmlquery v1.2.4
	github.com/bwmarrin/discordgo v0.23.3-0.20211010150959-f0b7e81468f7
	github.com/go-toast/toast v0.0.0-20190211030409-01e6764cf0a4
	github.com/joho/godotenv v1.4.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
)

require (
	github.com/antchfx/xpath v1.2.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	golang.org/x/text v0.3.3 // indirect
)
