package command

import (
	"io/ioutil"
	"net/http"
)

func getRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		ExitWithError(ExitBadConnection, err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
