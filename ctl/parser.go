package ctl

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/oZhuZhiYuan/kongctl/ctl/command"
)

func pflagParser(cfg *command.GlobalFlags) {
	if len(cfg.Hosts) == 1 && cfg.Hosts[0] == "127.0.0.1:8001" {
		// fmt.Println("hosts flag ok")
		return
	}
	if err := checkIPs(cfg); err != nil {
		command.ExitWithError(command.ExitBadFlag, err)
	}

}

func checkIPs(cfg *command.GlobalFlags) error {
	hosts := cfg.Hosts
	for index, host := range hosts {
		ipp := strings.Split(host, ":")
		if len(ipp) > 2 {
			return fmt.Errorf("flag hosts: host[:port]")
		}
		// check ip format
		if len(ipp) == 2 {
			port, err := strconv.Atoi(ipp[1])
			if err != nil {
				return fmt.Errorf("flag hosts: port must be integer")
			}
			if port < 0 || port > 65535 {
				return fmt.Errorf("flag hosts: port must > 0 and < 65535")
			}

			if err := checkHostFormat(ipp[0]); err != nil {
				return err
			}
		}
		if len(ipp) == 1 {
			if err := checkHostFormat(ipp[0]); err != nil {
				return err
			}
			cfg.Hosts[index] = ipp[0] + ":8001"
		}
		// fmt.Println("checkip", ip, len(ip))
	}

	return nil
}

func checkHostFormat(host string) error {
	// check domain format.   mind  \\.
	domain := "^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][-_.a-zA-Z0-9]{0,61}[a-zA-Z0-9]))\\.([a-zA-Z]{2,13}|[a-zA-Z0-9-]{2,30}.[a-zA-Z]{2,3})$"
	re := regexp.MustCompile(domain)
	if re.MatchString(host) {
		return nil
	}

	// check ip format. mind  \\.
	ip := "^([1-9][0-9]?|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.(([1-9]?[0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){2}([1-9][0-9]?|1[0-9][0-9]|2[0-4][0-9]|25[0-4])$"
	re = regexp.MustCompile(ip)
	if re.MatchString(host) {
		return nil
	}
	// if neither return error
	return fmt.Errorf("flag hosts: ip format is illegal")
}
