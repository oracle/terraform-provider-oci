// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MarketplaceListingPackageAgreementDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceListingPackageAgreement,
		Schema: map[string]*schema.Schema{
			"agreement_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Computed
			"author": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prompt": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readMarketplaceListingPackageAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingPackageAgreementDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()
	return tfresource.ReadResource(sync)
}

type MarketplaceListingPackageAgreementDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.GetAgreementResponse
}

func (s *MarketplaceListingPackageAgreementDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceListingPackageAgreementDataSourceCrud) Get() error {
	request := oci_marketplace.GetAgreementRequest{}

	if agreementId, ok := s.D.GetOkExists("agreement_id"); ok {
		tmp := agreementId.(string)
		request.AgreementId = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists("package_version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.GetAgreement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MarketplaceListingPackageAgreementDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("author", s.Res.Author)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContentUrl != nil {
		s.D.Set("content_url", *s.Res.ContentUrl)
	}

	if s.Res.Prompt != nil {
		s.D.Set("prompt", *s.Res.Prompt)
	}

	if s.Res.Signature != nil {
		s.D.Set("signature", *s.Res.Signature)
	}

	return nil
}
