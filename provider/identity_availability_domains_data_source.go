// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func AvailabilityDomainsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAvailabilityDomains,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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

func readAvailabilityDomains(d *schema.ResourceData, m interface{}) error {
	sync := &AvailabilityDomainsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type AvailabilityDomainsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListAvailabilityDomainsResponse
}

func (s *AvailabilityDomainsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AvailabilityDomainsDataSourceCrud) Get() error {
	request := oci_identity.ListAvailabilityDomainsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListAvailabilityDomains(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AvailabilityDomainsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		availabilityDomain := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Name != nil {
			availabilityDomain["name"] = *r.Name
		}

		resources = append(resources, availabilityDomain)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AvailabilityDomainsDataSource().Schema["availability_domains"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("availability_domains", resources); err != nil {
		panic(err)
	}

	return
}
