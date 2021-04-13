package command

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

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
	Data []service
}

func (this *services) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "host", "protocol", "port", "path",
		"retries", "write_timeout", "read_timeout", "connect_timeout"})
	for _, obj := range this.Data {
		table.Append([]string{obj.Id, obj.Name, obj.Host, obj.Protocol, fmt.Sprint(obj.Port), obj.Path,
			fmt.Sprint(obj.Retries), fmt.Sprint(obj.Write_timeout),
			fmt.Sprint(obj.Read_timeout), fmt.Sprint(obj.Connect_timeout),
		})
	}
	table.Render()
}
