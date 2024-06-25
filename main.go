//go:build !std
// +build !std

package main

import (
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
)

type input struct {
	Repository  string `json:"repository"`
	Remote      string `json:"remote"`
	Branch      string `json:"branch"`
	ApiKey      string `json:"api_key"`
	GithubToken string `json:"github_token"`
}

//export run
func run() int32 {
	url := "https://api.greptile.com/v2/repositories"

	in := input{}
	err := pdk.InputJSON(&in)
	if err != nil {
		pdk.SetError(fmt.Errorf("failed to parse input JSON: %v", err))
		return 1
	}

	if in.Repository == "" {
		pdk.SetError(fmt.Errorf("must specify repository"))
		return 1
	}

	if in.Remote == "" {
		in.Remote = "github"
	}

	if in.Branch == "" {
		in.Branch = "main"
	}

	payload := map[string]interface{}{
		"remote":     in.Remote,
		"repository": in.Repository,
		"branch":     in.Branch,
		"reload":     true,
		"notify":     true,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		pdk.SetError(fmt.Errorf("failed to marshal payload: %v", err))
		return 1
	}

	req := pdk.NewHTTPRequest(pdk.MethodPost, url)
	req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", in.ApiKey))
	req.SetHeader("X-GitHub-Token", in.GithubToken)
	req.SetHeader("Content-Type", "application/json")
	req.SetBody(payloadBytes)

	res := req.Send()

	if res.Status() != 200 {
		pdk.SetError(fmt.Errorf("failed to index repository: status %d", res.Status()))
		return 1
	}

	pdk.OutputString(fmt.Sprintf("Response status: %d\nBody: %s", res.Status(), string(res.Body())))
	return 0
}

func main() {}
