package scp

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var HelpCmd = &Command{
	Name:        "help",
	Description: " ",
	Options: []*ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "all",
			Description: "Shows a list of all commands",
			Handler:     AllHelp,
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "command",
			Description: "Shows info about a specified command",
			Options: []*ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "command",
					Description: "Command to get info about",
					Required:    true,
				},
			},
			Handler: CommandHelp,
		},
	},
}

func AllHelp(ctx *Ctx) {

}

func CommandHelp(ctx *Ctx) {
	ctx.Respond(5, &discordgo.InteractionResponseData{})
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("❓ Help: %v", ctx.Options[0].Value),
	}
	for _, c := range ctx.Router.AllCommands {
		if c.Name == ctx.Options[0].Value {

		}
	}
}

func (r *Router) UseHelpCommand(category *Category) {
	if category == nil {
		r.AddCommand(HelpCmd)
	} else {
		category.AddCommand(HelpCmd)
	}
	r.AllCommands = append(r.AllCommands, HelpCmd)
}
