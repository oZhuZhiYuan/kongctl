package command

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

// upstream
type upstream struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Create_at int64  `json:"created_at"`
}

type upstreams struct {
	Data []upstream `json:"data"`
}

func (this *upstreams) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "create_at"})
	for _, obj := range this.Data {
		if obj.Id == "" {
			continue
		}
		tm := time.Unix(obj.Create_at, 0)
		table.Append([]string{obj.Id, obj.Name,
			tm.Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

// target
type target struct {
	Id         string `json:"id"`
	Target     string `json:"target"`
	Weight     int    `json:"weight"`
	Upsteam    string
	Health     string  `json:"health"`
	Created_at float64 `json:"created_at"`
}

type targets struct {
	Data []target `json:"data"`
}

func (this *targets) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "target", "weight", "upstream", "health", "created_at"})
	for _, obj := range this.Data {
		if obj.Id == "" {
			continue
		}
		tm := time.Unix(int64(obj.Created_at), 0)
		table.Append([]string{obj.Id, obj.Target, fmt.Sprint(obj.Weight),
			obj.Upsteam, obj.Health, tm.Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

// targetResp
type targetResp struct {
	Id         string `json:"id"`
	Target     string `json:"target"`
	Weight     int    `json:"weight"`
	Upsteam    string
	Created_at float64 `json:"created_at"`
}

type targetResps struct {
	Data []targetResp `json:"data"`
}

func (this targetResps) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "target", "weight", "upstream"})
	for _, obj := range this.Data {
		if obj.Id == "" {
			continue
		}
		table.Append([]string{obj.Id, obj.Target, fmt.Sprint(obj.Weight),
			obj.Upsteam})
	}
	table.Render()
}

// targetPost
type targetPost struct {
	Target string `json:"target"`
	Weight int    `json:"weight"`
}
