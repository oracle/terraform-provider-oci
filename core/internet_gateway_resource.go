// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func InternetGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createInternetGateway,
		Read:     readInternetGateway,
		Update:   updateInternetGateway,
		Delete:   deleteInternetGateway,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": {
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

func createInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)

	return crud.CreateResource(d, sync)
}

func readInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)

	return crud.ReadResource(sync)
}

func updateInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)

	return crud.UpdateResource(d, sync)

}

func deleteInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}
