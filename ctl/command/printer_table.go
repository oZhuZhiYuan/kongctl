package command

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func upStreamsPrint(upss *upstreams) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "create_at"})
	for _, ups := range upss.Data {
		tm := time.Unix(ups.Create_at, 0)
		table.Append([]string{ups.Id, ups.Name,
			tm.Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

func targetsPrint(tgts *targets) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "target", "weight", "upstream", "health", "created_at"})
	for _, tgt := range tgts.Data {
		tm := time.Unix(int64(tgt.Created_at), 0)
		table.Append([]string{tgt.Id, tgt.Target, strconv.Itoa(tgt.Weight),
			tgt.Upsteam, tgt.Health, tm.Format("2006-01-02 15:04:05")})
	}
	table.Render()
}

func targetRespPrint(tgts []targetResp) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "target", "weight", "upstream"})
	for _, tgt := range tgts {
		table.Append([]string{tgt.Id, tgt.Target, strconv.Itoa(tgt.Weight),
			tgt.Upsteam})
	}
	table.Render()
}

func servicesPrint(objs *services) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "name", "host", "protocol", "port", "path",
		"retries", "write_timeout", "read_timeout", "connect_timeout"})
	for _, obj := range objs.Data {
		table.Append([]string{obj.Id, obj.Name, obj.Host, obj.Protocol, fmt.Sprint(obj.Port), obj.Path,
			fmt.Sprint(obj.Retries), fmt.Sprint(obj.Write_timeout),
			fmt.Sprint(obj.Read_timeout), fmt.Sprint(obj.Connect_timeout),
		})
	}
	table.Render()
}
