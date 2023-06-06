package association

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_ip_association", func(r *config.Resource) {

		r.ShortGroup = "ip"
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.OmittedFields = []string{
			"name",
		}
	})
}
