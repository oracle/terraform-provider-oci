// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreAppCatalogListingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreAppCatalogListings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publisher_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publisher_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_catalog_listings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"contact_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"listing_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"publisher_logo_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"publisher_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_published": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreAppCatalogListings(d *schema.ResourceData, m interface{}) error {
	sync := &CoreAppCatalogListingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreAppCatalogListingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListAppCatalogListingsResponse
}

func (s *CoreAppCatalogListingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreAppCatalogListingsDataSourceCrud) Get() error {
	request := oci_core.ListAppCatalogListingsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if publisherName, ok := s.D.GetOkExists("publisher_name"); ok {
		tmp := publisherName.(string)
		request.PublisherName = &tmp
	}

	if publisherType, ok := s.D.GetOkExists("publisher_type"); ok {
		tmp := publisherType.(string)
		request.PublisherType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListAppCatalogListings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAppCatalogListings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreAppCatalogListingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreAppCatalogListingsDataSource-", CoreAppCatalogListingsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		appCatalogListing := map[string]interface{}{}

		if r.DisplayName != nil {
			appCatalogListing["display_name"] = *r.DisplayName
		}

		if r.ListingId != nil {
			appCatalogListing["listing_id"] = *r.ListingId
		}

		if r.PublisherName != nil {
			appCatalogListing["publisher_name"] = *r.PublisherName
		}

		if r.Summary != nil {
			appCatalogListing["summary"] = *r.Summary
		}

		resources = append(resources, appCatalogListing)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreAppCatalogListingsDataSource().Schema["app_catalog_listings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("app_catalog_listings", resources); err != nil {
		return err
	}

	return nil
}
