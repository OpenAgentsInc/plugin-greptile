//go:build !std
// +build !std

package main

import (
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
)

type input struct {
	Operation   string `json:"operation"`
	Repository  string `json:"repository"`
	Remote      string `json:"remote"`
	Branch      string `json:"branch"`
	ApiKey      string `json:"api_key"`
	GithubToken string `json:"github_token"`
	Query       string `json:"query"`
	Messages    []struct {
		ID      string `json:"id"`
		Content string `json:"content"`
		Role    string `json:"role"`
	} `json:"messages"`
	SessionID string `json:"session_id"`
	Stream    bool   `json:"stream"`
	Genius    bool   `json:"genius"`
}

//export run
func run() int32 {
	in := input{}
	err := pdk.InputJSON(&in)
	if err != nil {
		pdk.SetError(fmt.Errorf("failed to parse input JSON: %v", err))
		return 1
	}

	switch in.Operation {
	case "index":
		return indexRepository(in)
	case "query":
		return queryRepository(in)
	case "search":
		return searchRepository(in)
	default:
		pdk.SetError(fmt.Errorf("invalid operation: %s", in.Operation))
		return 1
	}
}

func indexRepository(in input) int32 {
	url := "https://api.greptile.com/v2/repositories"

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

	return sendRequest(url, in.ApiKey, in.GithubToken, payload)
}

func queryRepository(in input) int32 {
	url := "https://api.greptile.com/v2/query"

	payload := map[string]interface{}{
		"messages": in.Messages,
		"repositories": []map[string]string{{
			"remote":     in.Remote,
			"branch":     in.Branch,
			"repository": in.Repository,
		}},
		"sessionId": in.SessionID,
		"stream":    in.Stream,
		"genius":    in.Genius,
	}

	return sendRequest(url, in.ApiKey, in.GithubToken, payload)
}

func searchRepository(in input) int32 {
	url := "https://api.greptile.com/v2/search"

	if in.Query == "" {
		pdk.SetError(fmt.Errorf("must specify a search query"))
		return 1
	}

	payload := map[string]interface{}{
		"query": in.Query,
		"repositories": []map[string]string{{
			"remote":     in.Remote,
			"branch":     in.Branch,
			"repository": in.Repository,
		}},
		"sessionId": in.SessionID,
		"stream":    in.Stream,
	}

	return sendRequest(url, in.ApiKey, in.GithubToken, payload)
}

func sendRequest(url, apiKey, githubToken string, payload interface{}) int32 {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		pdk.SetError(fmt.Errorf("failed to marshal payload: %v", err))
		return 1
	}

	req := pdk.NewHTTPRequest(pdk.MethodPost, url)
	req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.SetHeader("X-GitHub-Token", githubToken)
	req.SetHeader("Content-Type", "application/json")
	req.SetBody(payloadBytes)

	res := req.Send()

	if res.Status() != 200 {
		pdk.SetError(fmt.Errorf("request failed: status %d", res.Status()))
		return 1
	}

	pdk.OutputString(fmt.Sprintf("Response status: %d\nBody: %s", res.Status(), string(res.Body())))
	return 0
}

func main() {}
