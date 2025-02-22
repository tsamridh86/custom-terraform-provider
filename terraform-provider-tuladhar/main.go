package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/tsamridh86/terraform-provider-tuladhar/tuladhar"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: tuladhar.Provider,
	})
}
