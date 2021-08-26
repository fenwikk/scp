package scp

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Router struct {
	Cateories     []*Category
	Uncategorized []*Command
	AllCommands   []*Command
	Session       *discordgo.Session
}

func Create(s *discordgo.Session) *Router {
	return &Router{
		Session: s,
	}
}

func (r *Router) RegisterAllCommands(guildID string) {
	log.Println("Registering commands...")
	r.AllCommands = r.Uncategorized

	for _, c := range r.Cateories {
		c.RegisterCmds(c)
		for _, cmd := range c.Commands {
			r.AllCommands = append(r.AllCommands, cmd)
		}
	}

	for _, v := range r.AllCommands {
		_, err := r.Session.ApplicationCommandCreate(r.Session.State.User.ID, guildID, v.ToApplicationCommand())
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		log.Println("Created", v.Name)
	}

	r.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		guild, _ := s.Guild(i.GuildID)
		for _, command := range r.AllCommands {
			if command.Name == i.ApplicationCommandData().Name {
				log.Println("Running", command.Name, "in", guild.Name, "(", guild.ID, ")")
				guild, _ := s.Guild(i.GuildID)
				channel, _ := s.Channel(i.ChannelID)

				ctx := &Ctx{
					Options:     i.ApplicationCommandData().Options,
					Session:     s,
					Interaction: i,
					Router:      r,
					Guild:       guild,
					Channel:     channel,
					User:        i.Member.User,
				}

				for _, o := range command.Options {
					if i.ApplicationCommandData().Options[0].Name == o.Name && (o.Type == discordgo.ApplicationCommandOptionSubCommand || o.Type == discordgo.ApplicationCommandOptionSubCommandGroup) {
						o.Handler(ctx)
						return
					}
				}

				command.Handler(ctx)
			}
		}
	})
}

func (r *Router) AddCategory(c *Category) {
	r.Cateories = append(r.Cateories, c)

	c.RegisterCmds(c)
	for _, cmd := range c.Commands {
		r.AllCommands = append(r.AllCommands, cmd)
	}
}

func (r *Router) AddCommand(c *Command) {
	r.Uncategorized = append(r.Uncategorized, c)
}

func (r *Router) GetCategory(name string) *Category {
	for _, c := range r.Cateories {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (r *Router) GetCommand(name string) *Command {
	for _, c := range r.Uncategorized {
		if c.Name == name {
			return c
		}
	}
	return nil
}
