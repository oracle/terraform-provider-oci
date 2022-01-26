// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityRegionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityRegions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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

func readIdentityRegions(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityRegionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityRegionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListRegionsResponse
}

func (s *IdentityRegionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityRegionsDataSourceCrud) Get() error {
	request := oci_identity.ListRegionsRequest{}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListRegions(context.Background())
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityRegionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityRegionsDataSource-", IdentityRegionsDataSource(), s.D))
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityRegionsDataSource().Schema["regions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("regions", resources); err != nil {
		return err
	}

	return nil
}
