package command

import (
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
		serviceShowRoutes(),
		serviceShowPlugins(),
	)
	return cmd
}

// ########## kongctl service show
func serviceShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [--sname service-names]",
		Short: "Show services",
		Long: `show all services by default,
if set flag sname,show the provided services only.
		`,
		Run: serviceShowOp,
	}
	cmd.Flags().StringSliceVar(&snames, "sname", []string{}, "service names")
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
				getServiceByNname(host)
			}
		}
		return
	}
}

func getAllService(host string) {
	sers := services{}
	url := "http://" + host + "/services"
	getRequestJson(url, &sers)
	sers.printTable()
}

func getServiceByNname(host string) {
	sers := services{}
	sers.Data = make([]service, len(snames))
	for index, sname := range snames {
		url := "http://" + host + "/services/" + sname
		getRequestJson(url, &sers.Data[index])
	}
	sers.printTable()
}

// ########## kongctl service show-routes
func serviceShowRoutes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-routes [--sname service-names]",
		Short: "Show routes of services",
		Long:  `Show routes of services`,
		Run:   serviceShowRoutesOp,
	}
	cmd.Flags().StringSliceVar(&snames, "sname", []string{}, "service names")
	return cmd
}

func serviceShowRoutesOp(cmd *cobra.Command, args []string) {
	if len(snames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("show-routes need set service_name {--sname}"))
	}
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	banner := "Routes in host %s are as follow :\n"
	// Don't be concurret
	if concur == 0 {
		for index, host := range hosts {
			printBanner(fmt.Sprint(index+1), host, banner, len(hosts))
			getRoutes(host)
		}
		return
	}
}

func getRoutes(host string) {
	routes_all := routes{}
	for _, sname := range snames {
		url := "http://" + host + "/services/" + sname + "/routes"
		routes := routes{}
		getRequestJson(url, &routes)
		for i := 0; i < len(routes.Data); i++ {
			routes.Data[i].Service_name = sname
		}
		routes_all.Data = append(routes_all.Data, routes.Data...)
	}
	routes_all.printTable()
}

// ########## kongctl service show-plugins
func serviceShowPlugins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-plugins [--sname service-names]",
		Short: "Show plugins of services",
		Long:  `Show plugins of services`,
		Run:   serviceShowPluginsOp,
	}
	cmd.Flags().StringSliceVar(&snames, "sname", []string{}, "service names")
	return cmd
}

func serviceShowPluginsOp(cmd *cobra.Command, args []string) {
	if len(snames) == 0 {
		ExitWithError(ExitBadFlag, fmt.Errorf("show-plugins need set service_name {--sname}"))
	}
	hosts, _ := cmd.Flags().GetStringSlice("hosts")
	concur, _ := cmd.Flags().GetInt("concur")
	banner := "Plugins in host %s are as follow :\n"
	// Don't be concurret
	if concur == 0 {
		for index, host := range hosts {
			printBanner(fmt.Sprint(index+1), host, banner, len(hosts))
			getPlugins(host)
		}
		return
	}
}

func getPlugins(host string) {
	plugins_all := plugins{}
	for _, sname := range snames {
		url := "http://" + host + "/services/" + sname + "/plugins"
		plugins := plugins{}
		getRequestJson(url, &plugins)
		for i := 0; i < len(plugins.Data); i++ {
			plugins.Data[i].Service_name = sname
		}
		plugins_all.Data = append(plugins_all.Data, plugins.Data...)
	}
	plugins_all.printTable()
}
