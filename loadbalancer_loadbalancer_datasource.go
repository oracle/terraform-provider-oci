// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
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
	client := m.(client.BareMetalClient)
	sync := &LoadBalancerDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
