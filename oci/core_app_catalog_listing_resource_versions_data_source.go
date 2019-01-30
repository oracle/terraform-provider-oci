// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func AppCatalogListingResourceVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAppCatalogListingResourceVersions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"app_catalog_listing_resource_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"accessible_ports": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"allowed_actions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"available_regions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compatible_shapes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"listing_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"listing_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"listing_resource_version": {
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

func readAppCatalogListingResourceVersions(d *schema.ResourceData, m interface{}) error {
	sync := &AppCatalogListingResourceVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

type AppCatalogListingResourceVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListAppCatalogListingResourceVersionsResponse
}

func (s *AppCatalogListingResourceVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AppCatalogListingResourceVersionsDataSourceCrud) Get() error {
	request := oci_core.ListAppCatalogListingResourceVersionsRequest{}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListAppCatalogListingResourceVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAppCatalogListingResourceVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AppCatalogListingResourceVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		appCatalogListingResourceVersion := map[string]interface{}{
			"listing_id": *r.ListingId,
		}

		if r.ListingResourceId != nil {
			appCatalogListingResourceVersion["listing_resource_id"] = *r.ListingResourceId
		}

		if r.ListingResourceVersion != nil {
			appCatalogListingResourceVersion["listing_resource_version"] = *r.ListingResourceVersion
		}

		if r.TimePublished != nil {
			appCatalogListingResourceVersion["time_published"] = r.TimePublished.String()
		}

		resources = append(resources, appCatalogListingResourceVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AppCatalogListingResourceVersionsDataSource().Schema["app_catalog_listing_resource_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("app_catalog_listing_resource_versions", resources); err != nil {
		return err
	}

	return nil
}
