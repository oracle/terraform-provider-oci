// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerShapeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerShapes,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shapes": {
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

func readLoadBalancerShapes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &LoadBalancerShapeDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type LoadBalancerShapeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListLoadBalancerShapes
}

func (s *LoadBalancerShapeDatasourceCrud) Get() (e error) {
	cID := s.D.Get("compartment_id").(string)
	s.Res, e = s.Client.ListLoadBalancerShapes(cID, nil)
	return
}

func (s *LoadBalancerShapeDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Res.LoadBalancerShapes {
			res := map[string]interface{}{
				"name": v.Name,
			}
			resources = append(resources, res)

		}
		s.D.Set("shapes", resources)
	}
	return
}
