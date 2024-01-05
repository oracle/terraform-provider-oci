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

func MarketplaceAcceptedAgreementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMarketplaceAcceptedAgreement,
		Read:     readMarketplaceAcceptedAgreement,
		Update:   updateMarketplaceAcceptedAgreement,
		Delete:   deleteMarketplaceAcceptedAgreement,
		Schema: map[string]*schema.Schema{
			// Required
			"agreement_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
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
			"signature": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMarketplaceAcceptedAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceAcceptedAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.CreateResource(d, sync)
}

func readMarketplaceAcceptedAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceAcceptedAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

func updateMarketplaceAcceptedAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceAcceptedAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMarketplaceAcceptedAgreement(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceAcceptedAgreementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MarketplaceAcceptedAgreementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_marketplace.MarketplaceClient
	Res                    *oci_marketplace.AcceptedAgreement
	DisableNotFoundRetries bool
}

func (s *MarketplaceAcceptedAgreementResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MarketplaceAcceptedAgreementResourceCrud) Create() error {
	request := oci_marketplace.CreateAcceptedAgreementRequest{}

	if agreementId, ok := s.D.GetOkExists("agreement_id"); ok {
		tmp := agreementId.(string)
		request.AgreementId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	if packageVersion, ok := s.D.GetOkExists("package_version"); ok {
		tmp := packageVersion.(string)
		request.PackageVersion = &tmp
	}

	if signature, ok := s.D.GetOkExists("signature"); ok {
		tmp := signature.(string)
		request.Signature = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	response, err := s.Client.CreateAcceptedAgreement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AcceptedAgreement
	return nil
}

func (s *MarketplaceAcceptedAgreementResourceCrud) Get() error {
	request := oci_marketplace.GetAcceptedAgreementRequest{}

	tmp := s.D.Id()
	request.AcceptedAgreementId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	response, err := s.Client.GetAcceptedAgreement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AcceptedAgreement
	return nil
}

func (s *MarketplaceAcceptedAgreementResourceCrud) Update() error {
	request := oci_marketplace.UpdateAcceptedAgreementRequest{}

	tmp := s.D.Id()
	request.AcceptedAgreementId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	response, err := s.Client.UpdateAcceptedAgreement(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AcceptedAgreement
	return nil
}

func (s *MarketplaceAcceptedAgreementResourceCrud) Delete() error {
	request := oci_marketplace.DeleteAcceptedAgreementRequest{}

	tmp := s.D.Id()
	request.AcceptedAgreementId = &tmp

	if signature, ok := s.D.GetOkExists("signature"); ok {
		tmp := signature.(string)
		request.Signature = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "marketplace")

	_, err := s.Client.DeleteAcceptedAgreement(context.Background(), request)
	return err
}

func (s *MarketplaceAcceptedAgreementResourceCrud) SetData() error {
	if s.Res.AgreementId != nil {
		s.D.Set("agreement_id", *s.Res.AgreementId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
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
