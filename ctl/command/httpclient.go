package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getRequest(url string) []byte {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ExitWithError(ExitBadConnection, err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func postRequest(url string, data []byte) []byte {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		ExitWithError(ExitBadConnection, err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func getRequestJson(url string, obj interface{}) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ExitWithError(ExitBadConnection, err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Println("json.Unmarshal error")
		ExitWithError(ExitError, err)
	}
}
