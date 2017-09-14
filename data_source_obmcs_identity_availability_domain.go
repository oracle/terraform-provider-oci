// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func AvailabilityDomainDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readAvailabilityDomains,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readAvailabilityDomains(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &AvailabilityDomainDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type AvailabilityDomainDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListAvailabilityDomains
}

func (s *AvailabilityDomainDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListAvailabilityDomains(compartmentID)
	return
}

func (s *AvailabilityDomainDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.AvailabilityDomains {
			res := map[string]interface{}{
				"name":           v.Name,
				"compartment_id": v.CompartmentID,
			}
			resources = append(resources, res)
		}
		s.D.Set("availability_domains", resources)
	}
	return
}
