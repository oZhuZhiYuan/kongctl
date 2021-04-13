package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version string = "1.0.0"
)

func VersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the version of kongctl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("kongctl version:", version)
		},
	}
}
