// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IPSecConnectionStatusDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readIPSecDeviceStatus,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
	client := m.(*OracleClients)
	reader := &IPSecConnectionStatusDatasourceCrud{}
	reader.D = d
	reader.Client = client.client

	return crud.ReadResource(reader)
}

type IPSecConnectionStatusDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.IPSecConnectionDeviceStatus
}

func (s *IPSecConnectionStatusDatasourceCrud) Get() (e error) {
	ipsecID := s.D.Get("ipsec_id").(string)
	res, e := s.Client.GetIPSecConnectionDeviceStatus(ipsecID)
	if e == nil {
		s.Resource = res
	}
	return
}

func (s *IPSecConnectionStatusDatasourceCrud) SetData() {
	if s.Resource == nil {
		return
	}
	s.D.SetId(s.Resource.ID)
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("id", s.Resource.ID)
	s.D.Set("time_created", s.Resource.TimeCreated)

	resources := []map[string]interface{}{}

	for _, val := range s.Resource.Tunnels {
		tunnel := map[string]interface{}{
			"ip_address":         val.IPAddress,
			"state":              val.State,
			"time_created":       val.TimeCreated.String(),
			"time_state_modifed": val.TimeStateModified.String(),
		}

		resources = append(resources, tunnel)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("tunnels", resources); err != nil {
		panic(err)
	}
}
