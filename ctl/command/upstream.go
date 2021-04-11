package command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	unames []string
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
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
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
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamShowOp(cmd *cobra.Command, args []string) {
	// fmt.Println(cmd.Flags().GetStringSlice("uname"))
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	for index, host := range hosts {
		fmt.Printf("\033[1;36;40m[%d]\033[0m upstreams in host \033[1;33;40m%s\033[0m are as follow :\n", index+1, host)

		if len(unames) == 0 {
			getAllUpstream(host)
		}
		if len(unames) > 0 {
			getUpstreamByNname(host)
		}
	}

}

func getAllUpstream(host string) {
	url := "http://" + host + "/upstreams"
	// fmt.Println(url)
	body := getRequest(url)
	upss := upstreams{}
	err := json.Unmarshal(body, &upss)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	upStreamsPrint(&upss)

}

func getUpstreamByNname(host string) {
	upss := upstreams{}
	for _, uname := range unames {
		url := "http://" + host + "/upstreams/" + uname
		body := getRequest(url)
		ups := upstream{}
		err := json.Unmarshal(body, &ups)
		if err != nil {
			ExitWithError(ExitError, err)
		}
		// if get nothing, skip print
		if ups == (upstream{}) {
			continue
		}
		upss.Data = append(upss.Data, ups)
	}
	upStreamsPrint(&upss)
}
