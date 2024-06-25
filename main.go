package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/extism/go-pdk"
)

type input struct {
	Text         string `json:"text"`
	To           string `json:"to"`
	ApiKey       string `json:"api_key"`
	GithubToken  string `json:"github_token"`
}

//go:export run
func run() int32 {
	url := "https://api.greptile.com/v2/repositories"
	payload := strings.NewReader(`{
		"remote": "github",
		"repository": "openagentsinc/kb-wanix",
		"branch": "main",
		"reload": true,
		"notify": true
	}`)

	in := input{}
	err := pdk.InputJSON(&in)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	req.Header.Add("Authorization", "Bearer "+in.ApiKey)
	req.Header.Add("X-GitHub-Token", in.GithubToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	fmt.Println(res)
	fmt.Println(string(body))

	pdk.OutputString("Attempting to index repository")
	return 0
}

//go:export greet
func Greet() int32 {
	name := pdk.InputString()
	pdk.OutputString("Hello, " + name)
	return 0
}

func main() {
	fmt.Println("Hello, World!")
}
