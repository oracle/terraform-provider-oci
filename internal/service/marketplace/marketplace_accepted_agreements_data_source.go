// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MarketplaceAcceptedAgreementsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceAcceptedAgreements,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"accepted_agreement_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"package_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"accepted_agreements": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(MarketplaceAcceptedAgreementResource()),
			},
		},
	}
}

func readMarketplaceAcceptedAgreements(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceAcceptedAgreementsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplaceAcceptedAgreementsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListAcceptedAgreementsResponse
}

func (s *MarketplaceAcceptedAgreementsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceAcceptedAgreementsDataSourceCrud) Get() error {
	request := oci_marketplace.ListAcceptedAgreementsRequest{}

	if acceptedAgreementId, ok := s.D.GetOkExists("id"); ok {
		tmp := acceptedAgreementId.(string)
		request.AcceptedAgreementId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists("package_version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.ListAcceptedAgreements(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAcceptedAgreements(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplaceAcceptedAgreementsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MarketplaceAcceptedAgreementsDataSource-", MarketplaceAcceptedAgreementsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		acceptedAgreement := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AgreementId != nil {
			acceptedAgreement["agreement_id"] = *r.AgreementId
		}

		if r.DisplayName != nil {
			acceptedAgreement["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			acceptedAgreement["id"] = *r.Id
		}

		if r.ListingId != nil {
			acceptedAgreement["listing_id"] = *r.ListingId
		}

		if r.PackageVersion != nil {
			acceptedAgreement["package_version"] = *r.PackageVersion
		}

		if r.TimeAccepted != nil {
			acceptedAgreement["time_accepted"] = r.TimeAccepted.String()
		}

		resources = append(resources, acceptedAgreement)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MarketplaceAcceptedAgreementsDataSource().Schema["accepted_agreements"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("accepted_agreements", resources); err != nil {
		return err
	}

	return nil
}
