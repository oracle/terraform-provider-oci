// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancers,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerResource(),
			},
		},
	}
}

func readLoadBalancers(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &LoadBalancerDatasourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

type LoadBalancerDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancers
}

func (s *LoadBalancerDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancers(cID, nil)
	return
}

func (s *LoadBalancerDatasourceCrud) SetData() {
	if s.Res == nil {
		panic("LoadBalancer Resource is nil, cannot SetData")
	}
	s.D.SetId(time.Now().UTC().String())

	resources := make([]map[string]interface{}, len(s.Res.LoadBalancers))
	for i, v := range s.Res.LoadBalancers {
		ip_addresses := make([]string, len(v.IPAddresses))
		for i, ad := range v.IPAddresses {
			ip_addresses[i] = ad.IPAddress
		}
		resources[i] = map[string]interface{}{
			"id":             v.ID,
			"compartment_id": v.CompartmentID,
			"display_name":   v.DisplayName,
			"ip_addresses":   ip_addresses,
			"is_private":     v.IsPrivate,
			"shape":          v.Shape,
			"state":          v.State,
			"subnet_ids":     v.SubnetIDs,
			"time_created":   v.TimeCreated.String(),
		}

	}
	err := s.D.Set("load_balancers", resources)
	if err != nil {
		log.Printf("[ERROR] Failed to set load_balancers: %v", err)
	}
	return
}
