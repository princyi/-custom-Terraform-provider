package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/yourusername/terraform-provider-example/example"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: example.Provider,
	})
}
