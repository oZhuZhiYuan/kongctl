package command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	unames      []string
	targetNames []string
)

type cuptreams struct {
	upss upstreams
	host string
}

func UpStreamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upstream",
		Short: "Upstream command",
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

// ########## kongctl upstream show
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

func upStreamShowOp(cmd *cobra.Command, args []string) {
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	banner := "upstreams in host %s are as follow :\n"
	// Don't be concurret
	if concur == 0 {
		for index, host := range hosts {
			printBanner(fmt.Sprint(index+1), host, banner, len(hosts))
			if len(unames) == 0 {
				getAllUpstream(host)
			}
			if len(unames) > 0 {
				getUpstreamByNname(host)
			}
		}
		return
	}
}

func getAllUpstream(host string) {
	upss := upstreams{}
	url := "http://" + host + "/upstreams"
	getRequestJson(url, &upss)
	upss.printTable()
}

func getUpstreamByNname(host string) {
	objs := upstreams{}
	objs.Data = make([]upstream, len(unames))
	for index, uname := range unames {
		url := "http://" + host + "/upstreams/" + uname
		getRequestJson(url, &objs.Data[index])
	}
	objs.printTable()
}

// ########## kongctl upstream show-targets
func upStreamShowTargets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-targets",
		Short: "show the targets of upstreams",
		Run:   upStreamShowTargetsOp,
	}
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	return cmd
}

func upStreamShowTargetsOp(cmd *cobra.Command, args []string) {
	if len(unames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("show-targets need set upstream_name {--uname}"))
	}
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	banner := "Targets on host %s are as follow:\n"
	if concur == 0 {
		for index, host := range hosts {
			printBanner(fmt.Sprint(index+1), host, banner, len(hosts))
			getTargets(host)
		}
		return
	}
}

func getTargets(host string) {
	objs_all := targets{}
	for _, uname := range unames {
		url := "http://" + host + "/upstreams/" + uname + "/health/"
		objs := targets{}
		getRequestJson(url, &objs)
		for i := 0; i < len(objs.Data); i++ {
			objs.Data[i].Upsteam = uname
		}

		objs_all.Data = append(objs_all.Data, objs.Data...)
	}
	objs_all.printTable()
}

// ########## kongctl upstream add-targets
func upStreamAddTargets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-targets",
		Short: "add the targets of upstreams",
		Run:   upStreamAddTargetsOp,
	}
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	cmd.Flags().StringSliceVar(&targetNames, "target", []string{}, "target names")
	cmd.Flags().Int("weight", 100, "target weight, default 100")
	return cmd
}

func upStreamAddTargetsOp(cmd *cobra.Command, args []string) {
	if len(unames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("Add-targets need set upstream_name {--uname}"))
	}
	if len(targetNames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("Add-targets need set target {--target}"))
	}

	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	weight, _ := cmd.Flags().GetInt("weight")
	concur, _ := cmd.Flags().GetInt("concur")
	banner := "Add argets on host %s :\n"
	if concur == 0 {
		for index, host := range hosts {
			printBanner(fmt.Sprint(index+1), host, banner, len(hosts))
			addTargets(host, weight)
		}
		return
	}

}

func addTargets(host string, weight int) {
	// fmt.Println(host, targetNames, unames, weight)
	objs_all := targetResps{}
	for _, uname := range unames {
		url := "http://" + host + "/upstreams/" + uname + "/targets"
		for _, targetname := range targetNames {
			payload := targetPost{targetname, weight}
			js, err := json.Marshal(payload)
			// fmt.Printf(string(js))
			if err != nil {
				ExitWithError(ExitError, err)
			}
			body := postRequest(url, js)
			// fmt.Printf(string(body))
			obj := targetResp{}
			if err := json.Unmarshal(body, &obj); err != nil {
				fmt.Println("json.Unmarshal error")
				ExitWithError(ExitError, err)
			}
			obj.Upsteam = uname
			objs_all.Data = append(objs_all.Data, obj)
		}
	}
	objs_all.printTable()
}

// ########## kongctl upstream del-targets
func upStreamDelTargets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-targets",
		Short: "del the targets of upstreams",
		Run:   upStreamDelTargetsOp,
	}
	cmd.Flags().StringSliceVar(&unames, "uname", []string{}, "upstream names")
	cmd.Flags().StringSliceVar(&targetNames, "target", []string{}, "target names")
	return cmd
}

func upStreamDelTargetsOp(cmd *cobra.Command, args []string) {
	if len(unames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("Del-targets need set upstream_name {--uname}"))
	}
	if len(targetNames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("Del-targets need set target {--target}"))
	}

	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	banner := "Del argets on host %s :\n"
	// set a target's weight to zero equal to delete it
	weight := 0
	if concur == 0 {
		for index, host := range hosts {
			printBanner(fmt.Sprint(index+1), host, banner, len(hosts))
			addTargets(host, weight)
		}
		return
	}

}

// ########## kongctl upstream create
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
