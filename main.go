package main

import (
	"github.com/drmjo/terraform-provider-ca-fingerprint/finger"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: finger.Provider})
}
