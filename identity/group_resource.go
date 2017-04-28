// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

// ResourceIdentityGroup exposes an IdentityGroup Resource
func GroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createGroup,
		Read:     readGroup,
		Update:   updateGroup,
		Delete:   deleteGroup,
		Schema:   baseIdentitySchemaWithID,
	}
}

func createGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteGroup(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &GroupSync{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}
