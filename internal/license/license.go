// Package license gets licenses from the GitHub API.
package license

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

// Get gets license text, replacing placeholders with actual values.
func Get(spdxId string, ctx Context) (string, error) {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return "", err
	}
	response := struct {
		Body string
	}{}
	endpoint := fmt.Sprintf("licenses/%s", spdxId)
	err = client.Get(endpoint, &response)
	if err != nil {
		return "", err
	}
	return ctx.Render(response.Body), nil
}

// List gets all licenses. Licenses are represented as a map from their key to their
// name.
func List() (map[string]string, error) {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}
	response := []struct {
		Key  string
		Name string
	}{}
	const endpoint string = "licenses"
	err = client.Get("licenses", &response)
	if err != nil {
		return nil, err
	}
	licenses := make(map[string]string, len(response))
	for _, res := range response {
		licenses[res.Key] = res.Name
	}
	return licenses, nil
}
