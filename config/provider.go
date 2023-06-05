/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/fire-ant/provider-infoblox-nios/config/dns/aaaarecord"
	"github.com/fire-ant/provider-infoblox-nios/config/dns/arecord"
	"github.com/fire-ant/provider-infoblox-nios/config/dns/cnamerecord"
	"github.com/fire-ant/provider-infoblox-nios/config/dns/mxrecord"
	"github.com/fire-ant/provider-infoblox-nios/config/dns/ptrrecord"
	"github.com/fire-ant/provider-infoblox-nios/config/dns/srvrecord"
	"github.com/fire-ant/provider-infoblox-nios/config/dns/txtrecord"
	"github.com/fire-ant/provider-infoblox-nios/config/ip/allocation"
	"github.com/fire-ant/provider-infoblox-nios/config/ip/association"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv4/ipv4allocation"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv4/ipv4association"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv4/ipv4network"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv4/ipv4networkcontainer"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv6/ipv6allocation"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv6/ipv6association"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv6/ipv6network"
	"github.com/fire-ant/provider-infoblox-nios/config/ipv6/ipv6networkcontainer"
	"github.com/fire-ant/provider-infoblox-nios/config/network/view"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "infoblox-nios"
	modulePath     = "github.com/fire-ant/provider-infoblox-nios"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		allocation.Configure,
		association.Configure,
		ipv4allocation.Configure,
		ipv4association.Configure,
		ipv4network.Configure,
		ipv4networkcontainer.Configure,
		ipv6allocation.Configure,
		ipv6association.Configure,
		ipv6network.Configure,
		ipv6networkcontainer.Configure,
		arecord.Configure,
		aaaarecord.Configure,
		cnamerecord.Configure,
		mxrecord.Configure,
		ptrrecord.Configure,
		srvrecord.Configure,
		txtrecord.Configure,
		view.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
