// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func UserGroupMembershipResource() *schema.Resource {
	return &schema.Resource{
		Create: createUserGroupMembership,
		Read:   readUserGroupMembership,
		Delete: deleteUserGroupMembership,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createUserGroupMembership(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readUserGroupMembership(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func deleteUserGroupMembership(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}
