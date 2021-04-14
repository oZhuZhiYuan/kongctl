package command

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// service
type service struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Host            string `json:"host"`
	Protocol        string `json:"protocol"`
	Port            int    `json:"port"`
	Path            string `json:"path"`
	Retries         int    `json:"retries"`
	Write_timeout   int    `json:"write_timeout"`
	Read_timeout    int    `json:"read_timeout"`
	Connect_timeout int    `json:"connect_timeout"`
	Create_at       int64  `json:"created_at"`
}

type services struct {
	Data []service `json:"data"`
}

func (this *services) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "host", "protocol", "port", "path",
		"retries", "write_timeout", "read_timeout", "connect_timeout"})
	for _, obj := range this.Data {
		if obj.Id == "" {
			continue
		}
		table.Append([]string{obj.Id, obj.Name, obj.Host, obj.Protocol, fmt.Sprint(obj.Port), obj.Path,
			fmt.Sprint(obj.Retries), fmt.Sprint(obj.Write_timeout),
			fmt.Sprint(obj.Read_timeout), fmt.Sprint(obj.Connect_timeout),
		})
	}
	table.Render()
}

// route
type route struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Methods      []string `json:"methods"`
	Host         []string `json:"hosts"`
	Path         []string `json:"paths"`
	Service_name string
	Create_at    int64 `json:"created_at"`
}

type routes struct {
	Data []route `json:"data"`
}

func (this *routes) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "methods", "host", "paths", "service_name"})
	for _, obj := range this.Data {
		if obj.Id == "" {
			continue
		}
		table.Append([]string{obj.Id, obj.Name, fmt.Sprint(obj.Methods),
			fmt.Sprint(obj.Host), fmt.Sprint(obj.Path), obj.Service_name})
	}
	table.Render()
}

// plugins

type plugin struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Enabled      bool     `json:"enabled"`
	Protocol     []string `json:"protocols"`
	Run_on       string   `json:"run_on"`
	Service_name string
	Create_at    int64 `json:"created_at"`
}

type plugins struct {
	Data []plugin `json:"data"`
}

func (this *plugins) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "enabled", "protocols", "run_on", "service_name"})
	for _, obj := range this.Data {
		if obj.Id == "" {
			continue
		}
		table.Append([]string{obj.Id, obj.Name, fmt.Sprint(obj.Enabled),
			fmt.Sprint(obj.Protocol), obj.Run_on, obj.Service_name})
	}
	table.Render()
}
