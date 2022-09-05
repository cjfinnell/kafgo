package cmd

import (
	"github.com/cjfinnell/kafgo/internal"
	"github.com/spf13/cobra"
)

var consumeCmd = &cobra.Command{
	Use:     "consume [topic]",
	Aliases: []string{"c"},
	Short:   "Consume from specified topic",
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.Consume(
			cmd.Context(),
			cmd.Flag(bootstrapFlag).Value.String(),
			args[0],
			cmd.Flag(usernameFlag).Value.String(),
			cmd.Flag(passwordFlag).Value.String(),
		)
	},
}
