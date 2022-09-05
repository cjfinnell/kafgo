package cmd

import (
	"github.com/cjfinnell/kafgo/internal"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:     "ping",
	Aliases: []string{"p"},
	Short:   "Test connection to kafka cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag(saslFlag).Value.String() == "true" {
			return internal.PingSASL(
				cmd.Context(),
				cmd.Flag(bootstrapFlag).Value.String(),
				cmd.Flag(usernameFlag).Value.String(),
				cmd.Flag(passwordFlag).Value.String(),
			)
		}

		return internal.Ping(cmd.Context(), cmd.Flag(bootstrapFlag).Value.String())
	},
}
