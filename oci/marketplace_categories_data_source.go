// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v39/marketplace"
)

func init() {
	RegisterDatasource("oci_marketplace_categories", MarketplaceCategoriesDataSource())
}

func MarketplaceCategoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceCategories,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"categories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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

func readMarketplaceCategories(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceCategoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).marketplaceClient()

	return ReadResource(sync)
}

type MarketplaceCategoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListCategoriesResponse
}

func (s *MarketplaceCategoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceCategoriesDataSourceCrud) Get() error {
	request := oci_marketplace.ListCategoriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "marketplace")

	response, err := s.Client.ListCategories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCategories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplaceCategoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("MarketplaceCategoriesDataSource-", MarketplaceCategoriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		category := map[string]interface{}{}

		if r.Name != nil {
			category["name"] = *r.Name
		}

		resources = append(resources, category)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MarketplaceCategoriesDataSource().Schema["categories"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("categories", resources); err != nil {
		return err
	}

	return nil
}
