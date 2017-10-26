// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ProtocolDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readProtocols,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocols": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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

func readProtocols(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &ProtocolDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type ProtocolDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancerProtocols
}

func (s *ProtocolDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancerProtocols(cID, nil)
	return
}

func (s *ProtocolDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Res.LoadBalancerProtocols {
			res := map[string]interface{}{
				"name": v.Name,
			}
			resources = append(resources, res)

		}

		if f, fOk := s.D.GetOk("filter"); fOk {
			resources = ApplyFilters(f.(*schema.Set), resources)
		}

		s.D.Set("protocols", resources)
	}
	return
}
