package renderers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cvescan/cvescan/renderers/html"
	"github.com/cvescan/cvescan/renderers/json"
	"github.com/cvescan/cvescan/renderers/stdout"
)

func TestNewProvider(t *testing.T) {
	renderer, err := NewRenderer("stdout")
	assert.NoError(t, err)
	assert.IsType(t, stdout.Renderer{}, renderer)

	renderer, err = NewRenderer("json")
	assert.NoError(t, err)
	assert.IsType(t, json.Renderer{}, renderer)

	renderer, err = NewRenderer("html")
	assert.NoError(t, err)
	assert.IsType(t, html.Renderer{}, renderer)

	_, err = NewRenderer("test")
	assert.Error(t, err)
}
