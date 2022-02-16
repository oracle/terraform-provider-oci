// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v58/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MarketplaceListingPackageAgreementResource() *schema.Resource {
	return &schema.Resource{
		Create: createMarketplaceListingPackageAgreement,
		Read:   readMarketplaceListingPackageAgreement,
		Delete: deleteMarketplaceListingPackageAgreement,
		Schema: map[string]*schema.Schema{
			"agreement_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func createMarketplaceListingPackageAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingPackageAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()
	return tfresource.CreateResource(d, sync)
}

func readMarketplaceListingPackageAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingPackageAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()
	return tfresource.ReadResource(sync)
}

func deleteMarketplaceListingPackageAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingPackageAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()
	return tfresource.DeleteResource(d, sync)
}

type MarketplaceListingPackageAgreementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_marketplace.MarketplaceClient
	Res                    *oci_marketplace.GetAgreementResponse
	DisableNotFoundRetries bool
}

func (s *MarketplaceListingPackageAgreementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MarketplaceListingPackageAgreementResourceCrud) Create() error {
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

func (s *MarketplaceListingPackageAgreementResourceCrud) Get() error {
	return nil
}

func (s *MarketplaceListingPackageAgreementResourceCrud) Delete() error {
	return nil
}

func (s *MarketplaceListingPackageAgreementResourceCrud) SetData() error {
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
