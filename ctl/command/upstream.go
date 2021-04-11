package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	uname []string
)

func UpStreamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upstream",
		Short: "upstream command",
	}
	cmd.AddCommand(
		upStreamShow(),
		upStreamShowTargets(),
	)
	return cmd
}

func upStreamShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [--uname upstream-names]",
		Short: "show upstreams",
		Long: `show all upstreams by default,
if set flag uname,show the provided upstreams only.
		`,
		Run: upStreamShowOp,
	}
	cmd.Flags().StringSliceVar(&uname, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamShowTargets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-targets {}",
		Short: "show the targets of upstreams",
		Run: func(cmd *cobra.Command, args []string) {
			uname, err := cmd.Flags().GetStringSlice("uname")
			if err != nil {
				ExitWithError(ExitError, err)
			}
			if len(uname) == 0 {
				ExitWithError(ExitBadFlag, fmt.Errorf("show-targets need set upstream_name {--uname}"))
			}
			fmt.Println("pending ...")
		},
	}
	cmd.Flags().StringSliceVar(&uname, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamShowOp(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Flags().GetStringSlice("uname"))
	fmt.Println(cmd.Flags().GetStringSlice("hosts"))
}
