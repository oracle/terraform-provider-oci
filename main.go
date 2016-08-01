package main

import (
	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	client := &baremtlsdk.Client{}
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider(client)
		},
	})
}
