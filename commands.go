package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Type: discordgo.MessageApplicationCommand,
		Name: "translate",
	},
}

func RegisterCommands(dg *discordgo.Session, guildID string) {
	for _, v := range commands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
	dg.AddHandler(commandHandler)
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}
	data := i.ApplicationCommandData()

	switch data.Name {
	case "translate":
		handleTranslate(s, i)
	}

}
func handleTranslate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         "test",
			Components:      nil,
			Embeds:          nil,
			AllowedMentions: nil,
			Flags:           uint64(discordgo.MessageFlagsEphemeral),
			Files:           nil,
			Choices:         nil,
			CustomID:        "",
			Title:           "",
		},
	})
}
