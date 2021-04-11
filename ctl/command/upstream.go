package command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	unames []string
)

type cuptreams struct {
	upss upstreams
	host string
}

func UpStreamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upstream",
		Short: "upstream command",
	}
	cmd.AddCommand(
		upStreamShow(),
		upStreamShowTargets(),
		upStreamAddTargets(),
		upStreamDelTargets(),
		upStreamCreate(),
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
		Use:   "show-targets",
		Short: "show the targets of upstreams",
		Run:   upStreamShowTargetsOp,
	}
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamAddTargets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-targets",
		Short: "add the targets of upstreams",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pending...")
		},
	}
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamDelTargets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-targets",
		Short: "del the targets of upstreams",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pending...")
		},
	}
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create upstreams",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pending...")
		},
	}
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamShowOp(cmd *cobra.Command, args []string) {
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	// Don't be concurret
	if concur == 0 {
		for index, host := range hosts {
			fmt.Printf("\033[1;36;40m[%d]\033[0m upstreams in host \033[1;33;40m%s\033[0m are as follow :\n", index+1, host)

			if len(unames) == 0 {
				getAllUpstream(host)
			}
			if len(unames) > 0 {
				getUpstreamByNname(host)
			}
		}
		return
	}
	// Be concurret
	cu := make(chan cuptreams, concur)
	for _, host := range hosts {
		go getUpstream(host, cu)
	}
	index := 1
	for cup := range cu {
		fmt.Printf("\033[1;36;40m[%d]\033[0m upstreams in host \033[1;33;40m%s\033[0m are as follow :\n", index, cup.host)
		upStreamsPrint(&cup.upss)
		index++
		if index > len(hosts) {
			break
		}
	}
}

func getUpstream(host string, cu chan cuptreams) {
	if len(unames) == 0 {
		getAllUpstreamC(host, cu)
	}
	if len(unames) > 0 {
		getUpstreamByNnameC(host, cu)
	}
}

func getAllUpstreamC(host string, cu chan cuptreams) {
	url := "http://" + host + "/upstreams"
	// fmt.Println(url)
	body := getRequest(url)
	upss := upstreams{}
	err := json.Unmarshal(body, &upss)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	cup := cuptreams{upss, host}
	cu <- cup

}

func getUpstreamByNnameC(host string, cu chan cuptreams) {
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
	cup := cuptreams{upss, host}
	cu <- cup
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

func upStreamShowTargetsOp(cmd *cobra.Command, args []string) {
	if len(unames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("show-targets need set upstream_name {--uname}"))
	}
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	if concur == 0 {
		for index, host := range hosts {
			fmt.Printf("\033[1;36;40m[%d]\033[0m upstreams in host \033[1;33;40m%s\033[0m are as follow :\n", index+1, host)
			getTargets(host)
		}
		return
	}
}

func getTargets(host string) {
	tgts_all := targets{}
	for _, uname := range unames {
		url := "http://" + host + "/upstreams/" + uname + "/health/"
		body := getRequest(url)
		tgts := targets{}
		if err := json.Unmarshal(body, &tgts); err != nil {
			ExitWithError(ExitError, err)
		}
		for i := 0; i < len(tgts.Data); i++ {
			tgts.Data[i].Upsteam = uname
		}

		tgts_all.Data = append(tgts_all.Data, tgts.Data...)
	}
	targetsPrint(&tgts_all)
}
