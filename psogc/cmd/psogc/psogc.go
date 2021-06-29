package psogc

import (
	"github.com/spf13/cobra"

	"github.com/msmsny/psogc/psogc/cmd/status"
)

func Execute() error {
	return NewPSOGCCommand().Execute()
}

func NewPSOGCCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:           "psogc",
		Short:         "Phantasy Start Online Episode I&II unofficial tools",
		Long:          "Phantasy Start Online Episode I&II unofficial tools",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmds.AddCommand(status.NewStatusCommand())

	return cmds
}
