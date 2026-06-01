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
