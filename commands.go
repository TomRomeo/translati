package main

import (
    "github.com/bwmarrin/discordgo"
    "github.com/DaikiYamakawa/deepl-go"
    "context"
    "fmt"
    "github.com/apex/log"
)

var deepL *deepl.Client

var commands = []*discordgo.ApplicationCommand{
    {
        Type: discordgo.MessageApplicationCommand,
        Name: "translate",
    },
}

func RegisterCommands(dg *discordgo.Session, guildID string) {

    // create a deepL client
    var err error
    deepL, err = deepl.New("https://api-free.deepl.com", nil)
    if err != nil {
        log.WithError(err).Error("failed to create deepL client")
        return
    }

    // log deepL information
    log.Info("deepL client created")
    log.Info("deepL client info:")
    ai, err := deepL.GetAccountStatus(context.Background())
    if err != nil {
        log.WithError(err).Error("failed to get deepL account info")
        return
    }
    log.Info(fmt.Sprintf("%+v", ai))


    // register the commands
    for _, v := range commands {
        _, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, v)
        if err != nil {
            log.WithError(err).Error("failed to register command: " + v.Name)
            return
        }
    }

    // add the command handler
    dg.AddHandler(commandHandler)
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

    // only react on application commands
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

    data := i.ApplicationCommandData()
    content := data.Resolved.Messages[data.TargetID].Content

    // translate the content
    translated, err := deepL.TranslateSentence(context.Background(), content, "", "EN")

    if err != nil {
        log.WithError(err).Error("failed to translate sentence")

        // send error message as an Ephemeral
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content:         "",
                Components:      nil,
                Embeds:          []*discordgo.MessageEmbed{
                    {
                        Title: "An Error occurred",
                        Description: err.Error(),
                        Color: 0xff0000,
                    },
                },
                AllowedMentions: nil,
                Flags:           uint64(discordgo.MessageFlagsEphemeral),
                Files:           nil,
                Choices:         nil,
                CustomID:        "",
                Title:           "",
            },
        })
        return
    }

    // send the translated text as an Ephemeral
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content:         "",
            Components:      nil,
            Embeds:          []*discordgo.MessageEmbed{
                {
                    Title: "Translated from: " + translated.Translations[0].DetectedSourceLanguage,
                    Description: "",
                    Color: 0x7289DA,
                    Fields: []*discordgo.MessageEmbedField{
                        {
                            Name: "Original",
                            Value: content,
                        },
                        {
                            Name: "Translated",
                            Value: translated.Translations[0].Text,
                        },
                    },
                    Footer: &discordgo.MessageEmbedFooter{
                        Text: "powered by deepl.com",
                    },
                },
            },
            AllowedMentions: nil,
            Flags:           uint64(discordgo.MessageFlagsEphemeral),
            Files:           nil,
            Choices:         nil,
            CustomID:        "",
            Title:           "",
        },
    })
}
