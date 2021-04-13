package command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	snames []string
)

func ServiceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service",
		Short: "Service command",
	}
	cmd.AddCommand(
		serviceShow(),
	)
	return cmd
}

func serviceShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [--sname upstream-names]",
		Short: "Show services",
		Long: `show all services by default,
if set flag sname,show the provided services only.
		`,
		Run: serviceShowOp,
	}
	cmd.Flags().StringSliceVar(&unames, "sname", []string{}, "service names")
	return cmd
}

func serviceShowOp(cmd *cobra.Command, args []string) {
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	banner := "services in host %s are as follow :\n"
	// Don't be concurret
	if concur == 0 {
		for index, host := range hosts {
			printBanner(fmt.Sprint(index+1), host, banner, len(hosts))
			if len(snames) == 0 {
				getAllService(host)
			}
			if len(snames) > 0 {
				fmt.Println(hosts)
			}
		}
		return
	}
}

func getAllService(host string) {
	sers := services{}
	url := "http://" + host + "/services"
	// fmt.Println(url)
	// body := getRequest(url)
	// err := json.Unmarshal(body, &sers)
	// if err != nil {
	// 	ExitWithError(ExitError, err)
	// }
	// servicesPrint(&sers)
	getObject(&sers, url)
}

func getServiceByNname(host string) {
	upss := upstreams{}
	for _, uname := range unames {
		url := "http://" + host + "/services/" + uname
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
