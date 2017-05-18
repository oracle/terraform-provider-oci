// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func VnicDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readVnic,
		Schema: map[string]*schema.Schema{
			"vnic_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname_label": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readVnic(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VnicDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type VnicDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.Vnic
}

func (v *VnicDatasourceCrud) Get() (e error) {
	id := v.D.Get("vnic_id").(string)

	v.Resource, e = v.Client.GetVnic(id)
	return
}

func (v *VnicDatasourceCrud) SetData() {
	if v.Resource != nil {
		v.D.SetId(v.Resource.ID)
		v.D.Set("id", v.Resource.ID)
		v.D.Set("availability_domain", v.Resource.AvailabilityDomain)
		v.D.Set("compartment_id", v.Resource.CompartmentID)
		v.D.Set("display_name", v.Resource.DisplayName)
		v.D.Set("hostname_label", v.Resource.HostnameLabel)
		v.D.Set("private_ip_address", v.Resource.PrivateIPAddress)
		v.D.Set("public_ip_address", v.Resource.PublicIPAddress)
		v.D.Set("state", v.Resource.State)
		v.D.Set("subnet_id", v.Resource.SubnetID)
	}
	return
}
