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

func MarketplacePublicationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["publication_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MarketplacePublicationResource(), fieldMap, readSingularMarketplacePublication)
}

func readSingularMarketplacePublication(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplacePublicationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.GetPublicationResponse
}

func (s *MarketplacePublicationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplacePublicationDataSourceCrud) Get() error {
	request := oci_marketplace.GetPublicationRequest{}

	if publicationId, ok := s.D.GetOkExists("publication_id"); ok {
		tmp := publicationId.(string)
		request.PublicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.GetPublication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MarketplacePublicationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Icon != nil {
		s.D.Set("icon", []interface{}{UploadDataToMap(s.Res.Icon)})
	} else {
		s.D.Set("icon", nil)
	}

	s.D.Set("listing_type", s.Res.ListingType)

	if s.Res.LongDescription != nil {
		s.D.Set("long_description", *s.Res.LongDescription)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("package_type", s.Res.PackageType)

	if s.Res.ShortDescription != nil {
		s.D.Set("short_description", *s.Res.ShortDescription)
	}

	s.D.Set("state", s.Res.LifecycleState)

	supportContacts := []interface{}{}
	for _, item := range s.Res.SupportContacts {
		supportContacts = append(supportContacts, SupportContactToMap(item))
	}
	s.D.Set("support_contacts", supportContacts)

	supportedOperatingSystems := []interface{}{}
	for _, item := range s.Res.SupportedOperatingSystems {
		supportedOperatingSystems = append(supportedOperatingSystems, OperatingSystemToMap(item))
	}
	s.D.Set("supported_operating_systems", supportedOperatingSystems)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
