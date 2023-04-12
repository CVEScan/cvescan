package snyk

import (
	"fmt"

	"github.com/kirinlabs/HttpRequest"

	"github.com/cvescan/cvescan/models"
)

const userAgent = "cvescan"

func newClient(c *models.Credentials) *HttpRequest.Request {
	return HttpRequest.NewRequest().SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("token %s", c.Token),
		"User-Agent":    userAgent,
	})
}
