// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/marketplace"
)

func init() {
	RegisterDatasource("oci_marketplace_accepted_agreement", MarketplaceAcceptedAgreementDataSource())
}

func MarketplaceAcceptedAgreementDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["accepted_agreement_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(MarketplaceAcceptedAgreementResource(), fieldMap, readSingularMarketplaceAcceptedAgreement)
}

func readSingularMarketplaceAcceptedAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceAcceptedAgreementDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).marketplaceClient()

	return ReadResource(sync)
}

type MarketplaceAcceptedAgreementDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.GetAcceptedAgreementResponse
}

func (s *MarketplaceAcceptedAgreementDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceAcceptedAgreementDataSourceCrud) Get() error {
	request := oci_marketplace.GetAcceptedAgreementRequest{}

	if acceptedAgreementId, ok := s.D.GetOkExists("accepted_agreement_id"); ok {
		tmp := acceptedAgreementId.(string)
		request.AcceptedAgreementId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "marketplace")

	response, err := s.Client.GetAcceptedAgreement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MarketplaceAcceptedAgreementDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AgreementId != nil {
		s.D.Set("agreement_id", *s.Res.AgreementId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ListingId != nil {
		s.D.Set("listing_id", *s.Res.ListingId)
	}

	if s.Res.PackageVersion != nil {
		s.D.Set("package_version", *s.Res.PackageVersion)
	}

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	return nil
}
