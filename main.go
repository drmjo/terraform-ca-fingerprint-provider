package main

import (
	"github.com/drmjo/terraform-provider-finger/finger"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: finger.Provider})
}
