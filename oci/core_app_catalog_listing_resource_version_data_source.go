// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func init() {
	RegisterDatasource("oci_core_app_catalog_listing_resource_version", CoreAppCatalogListingResourceVersionDataSource())
}

func CoreAppCatalogListingResourceVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreAppCatalogListingResourceVersion,
		Schema: map[string]*schema.Schema{
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_version": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularCoreAppCatalogListingResourceVersion(d *schema.ResourceData, m interface{}) error {
	sync := &CoreAppCatalogListingResourceVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

type CoreAppCatalogListingResourceVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetAppCatalogListingResourceVersionResponse
}

func (s *CoreAppCatalogListingResourceVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreAppCatalogListingResourceVersionDataSourceCrud) Get() error {
	request := oci_core.GetAppCatalogListingResourceVersionRequest{}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if resourceVersion, ok := s.D.GetOkExists("resource_version"); ok {
		tmp := resourceVersion.(string)
		request.ResourceVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetAppCatalogListingResourceVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreAppCatalogListingResourceVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("accessible_ports", s.Res.AccessiblePorts)

	s.D.Set("allowed_actions", s.Res.AllowedActions)

	s.D.Set("available_regions", s.Res.AvailableRegions)

	s.D.Set("compatible_shapes", s.Res.CompatibleShapes)

	if s.Res.ListingResourceId != nil {
		s.D.Set("listing_resource_id", *s.Res.ListingResourceId)
	}

	if s.Res.ListingResourceVersion != nil {
		s.D.Set("listing_resource_version", *s.Res.ListingResourceVersion)
	}

	if s.Res.TimePublished != nil {
		s.D.Set("time_published", s.Res.TimePublished.String())
	}

	return nil
}
