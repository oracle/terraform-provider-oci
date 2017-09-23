// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IPSecConnectionConfigDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readIPSecDeviceConfig,
		Schema: map[string]*schema.Schema{
			"ipsec_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shared_secret": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIPSecDeviceConfig(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	reader := &IPSecConnectionConfigDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type IPSecConnectionConfigDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.IPSecConnectionDeviceConfig
}

func (s *IPSecConnectionConfigDatasourceCrud) Get() (e error) {
	ipsecID := s.D.Get("ipsec_id").(string)
	res, e := s.Client.GetIPSecConnectionDeviceConfig(ipsecID)
	if e == nil {
		s.Resource = res
	}
	return
}

func (s *IPSecConnectionConfigDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(s.Resource.ID)
		s.D.Set("compartment_id", s.Resource.CompartmentID)
		s.D.Set("id", s.Resource.ID)
		s.D.Set("time_created", s.Resource.TimeCreated)

		tunnels := []map[string]interface{}{}

		for _, val := range s.Resource.Tunnels {
			tunnel := map[string]interface{}{
				"ip_address":    val.IPAddress,
				"shared_secret": val.SharedSecret,
				"time_created":  val.TimeCreated.String(),
			}

			tunnels = append(tunnels, tunnel)
		}

		s.D.Set("tunnels", tunnels)

	}
}
