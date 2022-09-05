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
		return internal.Ping(cmd.Context(), buildDialer(cmd), cmd.Flag(bootstrapFlag).Value.String())
	},
}
