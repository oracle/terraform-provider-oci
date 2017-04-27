// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func DBSystemResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: &crud.TwoHours,
			Delete: &crud.TwoHours,
			Update: &crud.TwoHours,
		},
		Create: createDBSystem,
		Read:   readDBSystem,
		Delete: deleteDBSystem,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_public_keys": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},
			"database_edition": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			"db_home": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"admin_password": {
										Type:     schema.TypeString,
										Required: true,
									},
									"db_name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"db_version": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"disk_redundancy": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"listener_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func deleteDBSystem(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}
