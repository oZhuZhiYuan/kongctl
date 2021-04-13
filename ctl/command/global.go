package command

import (
	"fmt"
	"strings"
)

type GlobalFlags struct {
	Hosts []string
}

func printBanner(index, host, banner string, num int) {
	hl_index := "\033[1;36;40m[" + index + "]\033[0m"
	hl_host := "\033[1;33;40m" + host + "\033[0m"
	// if run this on localhost, don't print index.
	if num == 1 && strings.HasPrefix(host, "127.0.0.1") {
		fmt.Printf(banner, hl_host)
	} else {
		banner = "%s " + banner
		fmt.Printf(banner, hl_index, hl_host)
	}
}
