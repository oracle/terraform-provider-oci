// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func RegionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readRegions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"regions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
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

func readRegions(d *schema.ResourceData, m interface{}) error {
	sync := &RegionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type RegionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListRegionsResponse
}

func (s *RegionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RegionsDataSourceCrud) Get() error {
	request := oci_identity.ListRegionsRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListRegions(context.Background())
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RegionsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		region := map[string]interface{}{}

		if r.Key != nil {
			region["key"] = *r.Key
		}

		if r.Name != nil {
			region["name"] = *r.Name
		}

		resources = append(resources, region)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, RegionsDataSource().Schema["regions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("regions", resources); err != nil {
		panic(err)
	}

	return
}
