package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider(client BareMetalClient) terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap:  resourcesMap(),
		ConfigureFunc: providerConfigure(client),
	}
}

func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_server":        ResourceServer(),
		"baremetal_identity_user": ResourceIdentityUser(),
	}
}

func providerConfigure(client BareMetalClient) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		return client, nil
	}
}
