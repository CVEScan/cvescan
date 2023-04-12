// Package providers contains functionality to retrieve vulnerability data from different external sources
package providers

import (
	"fmt"

	"github.com/cvescan/cvescan/models"
	"github.com/cvescan/cvescan/providers/ossindex"
	"github.com/cvescan/cvescan/providers/osv"
	"github.com/cvescan/cvescan/providers/snyk"
)

// NewProvider will return a provider interface for the requested vulnerability services
func NewProvider(name string) (provider models.Provider, err error) {
	switch name {
	case "ossindex":
		provider = ossindex.Provider{}
	case "osv":
		provider = osv.Provider{}
	case "snyk":
		provider = snyk.Provider{}
	default:
		err = fmt.Errorf("%s is not a valid provider type", name)
	}
	return
}
