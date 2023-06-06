package ipv4network

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_ipv4_network", func(r *config.Resource) {

		r.ShortGroup = "ipv4"
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.OmittedFields = []string{
			"name",
		}
	})
}
