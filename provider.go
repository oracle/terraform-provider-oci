package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider implementation for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: resourcesMap(),
	}
}


func resourcesMap() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"baremetal_server":        ResourceServer(),
		"baremetal_identity_user": ResourceIdentityUser(),
	}
}
