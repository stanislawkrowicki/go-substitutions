# go-substitutions
Go-substitutions listens for new substitutions on school's webpage

The notification (toast) service is designed to run on Windows 10 only,
however, the discord-service can run locally on every OS, or deployed to Heroku.

To run the app, you need an env file. It should be filled like this:
```dotenv
CLASS=YOUR_CLASS
SUBSTITUTIONS_API=API_URL
DISCORD_TOKEN=YOUR_BOT_TOKEN
DISCORD_GUILD=YOUR_DISCORD_GUILD
DISCORD_CHANNEL_ID=CHANNEL_TO_SEND_CHANGES
```

I didn't put the API URL in the code for security reasons.

Go-substitutions runs 24/7, so it's a good idea to have it run without terminal. To achieve this, build the app with:
```
go build -ldflags "-H windowsgui" cmd/toast-service/main.go
```

Then you can for example add a shortcut to the binary into your shell:startup, so it starts every time you turn on your PC.