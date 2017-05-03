// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func LoadBalancerPolicyDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readPolicies,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policies": {
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

func readPolicies(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &PoliciesDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type PoliciesDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancerPolicies
}

func (s *PoliciesDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancerPolicies(cID, nil)
	return
}

func (s *PoliciesDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Res.LoadBalancerPolicies {
			res := map[string]interface{}{
				"name": v.Name,
			}
			resources = append(resources, res)

		}
		s.D.Set("policies", resources)
	}
	return
}
