// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	provider "github.com/terraform-providers/terraform-provider-oci/oci"
)

func main() {
	provider.PrintVersion()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.Provider(provider.ProviderConfig)
		},
	})
}
