package ctl

import (
	"fmt"
	"os"

	"github.com/oZhuZhiYuan/kongctl/ctl/command"
	"github.com/spf13/cobra"
)

const (
	cliName        = "kongctl"
	cliDescription = "Command line for kong"
)

var (
	rootCmd = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"kongctl"},
	}
)

var (
	globalFlags = command.GlobalFlags{}
)

func init() {
	cobra.OnInitialize(initParser)
	rootCmd.PersistentFlags().StringSliceVarP(&globalFlags.Hosts, "hosts", "H", []string{"127.0.0.1:8001"}, "remote or local hosts")
	rootCmd.PersistentFlags().IntP("concur", "c", 0, "the number of concurrent opertion")
	rootCmd.AddCommand(
		command.UpStreamCommand(),
		command.ServiceCommand(),
		command.VersionCommand(),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initParser() {
	// fmt.Println("initParser", globalFlags.Hosts)
	pflagParser(&globalFlags)
}
