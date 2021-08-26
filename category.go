package scp

type Category struct {
	Name         string
	ID           string
	Description  string
	HelpEmoji    string
	Commands     []*Command
	RegisterCmds func(c *Category)
}

func (c *Category) AddCommand(cmd *Command) {
	c.Commands = append(c.Commands, cmd)
}

func (c *Category) GetCommand(name string) *Command {
	for _, cmd := range c.Commands {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}
