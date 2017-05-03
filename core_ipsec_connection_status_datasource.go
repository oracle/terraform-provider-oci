// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func IPSecConnectionStatusDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readIPSecDeviceStatus,
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
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_state_modifed": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIPSecDeviceStatus(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &IPSecConnectionStatusDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type IPSecConnectionStatusDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.IPSecConnectionDeviceStatus
}

func (s *IPSecConnectionStatusDatasourceCrud) Get() (e error) {
	ipsecID := s.D.Get("ipsec_id").(string)
	s.Resource, e = s.Client.GetIPSecConnectionDeviceStatus(ipsecID)
	return
}

func (s *IPSecConnectionStatusDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(s.Resource.ID)
		s.D.Set("compartment_id", s.Resource.CompartmentID)
		s.D.Set("id", s.Resource.ID)
		s.D.Set("time_created", s.Resource.TimeCreated)

		tunnels := []map[string]interface{}{}

		for _, val := range s.Resource.Tunnels {
			tunnel := map[string]interface{}{
				"ip_address":         val.IPAddress,
				"state":              val.State,
				"time_created":       val.TimeCreated.String(),
				"time_state_modifed": val.TimeStateModified.String(),
			}

			tunnels = append(tunnels, tunnel)
		}

		s.D.Set("tunnels", tunnels)

	}
}
