// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v32/identity"
)

func init() {
	RegisterDatasource("oci_identity_fault_domains", IdentityFaultDomainsDataSource())
}

func IdentityFaultDomainsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityFaultDomains,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fault_domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
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

func readIdentityFaultDomains(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityFaultDomainsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityFaultDomainsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListFaultDomainsResponse
}

func (s *IdentityFaultDomainsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityFaultDomainsDataSourceCrud) Get() error {
	request := oci_identity.ListFaultDomainsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListFaultDomains(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityFaultDomainsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("IdentityFaultDomainsDataSource-", IdentityFaultDomainsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		faultDomain := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.Id != nil {
			faultDomain["id"] = *r.Id
		}

		if r.Name != nil {
			faultDomain["name"] = *r.Name
		}

		resources = append(resources, faultDomain)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityFaultDomainsDataSource().Schema["fault_domains"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("fault_domains", resources); err != nil {
		return err
	}

	return nil
}
