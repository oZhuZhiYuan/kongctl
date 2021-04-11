package command

import (
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
	table.SetHeader([]string{"id", "target", "stream", "health", "created_at"})
	for _, tgt := range tgts.Data {
		tm := time.Unix(int64(tgt.Created_at), 0)
		table.Append([]string{tgt.Id, tgt.Target, strconv.Itoa(tgt.Weight),
			tgt.Upsteam, tgt.Health, tm.Format("2006-01-02 15:04:05")})
	}
	table.Render()
}
