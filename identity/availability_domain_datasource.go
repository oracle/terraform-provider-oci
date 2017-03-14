// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
	"github.com/hashicorp/terraform/helper/schema"
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
				Elem:     &schema.Resource{
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
	client := m.(client.BareMetalClient)
	sync := &AvailabilityDomainDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
