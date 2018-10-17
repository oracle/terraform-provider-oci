// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func AppCatalogListingDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAppCatalogListing,
		Schema: map[string]*schema.Schema{
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularAppCatalogListing(d *schema.ResourceData, m interface{}) error {
	sync := &AppCatalogListingDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return ReadResource(sync)
}

type AppCatalogListingDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetAppCatalogListingResponse
}

func (s *AppCatalogListingDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AppCatalogListingDataSourceCrud) Get() error {
	request := oci_core.GetAppCatalogListingRequest{}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetAppCatalogListing(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AppCatalogListingDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.ContactUrl != nil {
		s.D.Set("contact_url", *s.Res.ContactUrl)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.PublisherLogoUrl != nil {
		s.D.Set("publisher_logo_url", *s.Res.PublisherLogoUrl)
	}

	if s.Res.PublisherName != nil {
		s.D.Set("publisher_name", *s.Res.PublisherName)
	}

	if s.Res.Summary != nil {
		s.D.Set("summary", *s.Res.Summary)
	}

	if s.Res.TimePublished != nil {
		s.D.Set("time_published", s.Res.TimePublished.String())
	}

	return nil
}
