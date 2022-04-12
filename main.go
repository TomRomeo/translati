package main

import (
	"github.com/apex/log"
    "github.com/bwmarrin/discordgo"
    "github.com/joho/godotenv"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    // load .env file
    err := godotenv.Load()
    if err != nil {
        log.WithError(err).Error("Could ot load .env file")
    }

    // create discord struct
    dg, err := discordgo.New("Bot " + os.Getenv("BOT_KEY"))
    if err != nil {
        log.WithError(err).Error("An error occurred while trying to create the bot struct")
        return
    }

    // open discord session
    if err = dg.Open(); err != nil {
        log.WithError(err).Error("Could not establish a connection with Discord")
        return
    }
    log.Info("Successfully established a discord ws connection..")

    //register commands
	//RegisterCommands(dg, "862337011264126996")
	RegisterCommands(dg, "")
    log.Info("Registered command successfully...")

    log.Info("Ready to translate messages...")

    // graceful exit
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

    <-c
    log.Info("Shutting down...")
    if err = dg.Close(); err != nil {
        log.WithError(err).Error("Failed to close Discord connection properly")
        return
    }
}
