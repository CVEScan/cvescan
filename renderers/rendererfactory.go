// Package renderers contains functionality to render output in various formats
package renderers

import (
	"fmt"

	"github.com/cvescan/cvescan/models"
	"github.com/cvescan/cvescan/renderers/html"
	"github.com/cvescan/cvescan/renderers/json"
	"github.com/cvescan/cvescan/renderers/stdout"
)

// NewRenderer will return a Renderer interface for the requested output
func NewRenderer(output string) (renderer models.Renderer, err error) {
	switch output {
	case "stdout":
		renderer = stdout.Renderer{}
	case "json":
		renderer = json.Renderer{}
	case "html":
		renderer = html.Renderer{}
	default:
		err = fmt.Errorf("%s is not a valid output type", output)
	}
	return
}
