package main

import (
	"fmt"
	"os"
	"io"
	"encoding/json"
	"strings"
	"net/http"
)

const (
	// where we make GET requests for our CRON expressions
	base string = "https://cronprompt.com/api/prompt?prompt="
	usage string = "cwon [expression: string]. Example: cwon At midnight on Christmas"
)

type response struct {
	Cron string `json:"result"`
	Call string `json:"call"`
}

func main() {
	// shift the args to remove the program
	prompt := strings.Join(os.Args[1:], " ")
	if len(prompt) == 0 {
		fmt.Println("No expression provided.")
		fmt.Println(usage)
		os.Exit(0)
	}
	// construct the URL to which we will make our GET request
	url := base + prompt

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body into a byte array so we can unmarshal it later
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON response into a response struct
	data := response{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	fmt.Println(data.Cron)
}