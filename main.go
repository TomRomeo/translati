package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("BOT_KEY"))
	if err != nil {
		log.Fatalf("An error occurred while trying to create the bot:\n%s", err)
	}
	if err = dg.Open(); err != nil {
		log.Fatalf("Could not establish a connection with Discord:\n%s", err)
	}
	log.Println("Successfully established a discord ws connection..")

	//register commands
	RegisterCommands(dg, "862337011264126996")
	log.Println("Registered command successfully...")

	log.Println("Ready to translate messages...")

	// graceful exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-c
	log.Println("Shutting down...")
	if err = dg.Close(); err != nil {
		log.Println("Failed to close Discord connection properly")
	}
}
