// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v56/marketplace"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MarketplacePublicationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplacePublications,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"listing_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"operating_systems": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"publication_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(MarketplacePublicationResource()),
			},
		},
	}
}

func readMarketplacePublications(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublicationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplacePublicationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListPublicationsResponse
}

func (s *MarketplacePublicationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplacePublicationsDataSourceCrud) Get() error {
	request := oci_marketplace.ListPublicationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if listingType, ok := s.D.GetOkExists("listing_type"); ok {
		request.ListingType = oci_marketplace.ListPublicationsListingTypeEnum(listingType.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		interfaces := name.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("name") {
			request.Name = tmp
		}
	}

	if operatingSystems, ok := s.D.GetOkExists("operating_systems"); ok {
		interfaces := operatingSystems.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("operating_systems") {
			request.OperatingSystems = tmp
		}
	}

	if publicationId, ok := s.D.GetOkExists("id"); ok {
		tmp := publicationId.(string)
		request.PublicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.ListPublications(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPublications(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplacePublicationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MarketplacePublicationsDataSource-", MarketplacePublicationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		publication := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"listing_type":   r.ListingType,
		}

		if r.Icon != nil {
			publication["icon"] = []interface{}{UploadDataToMap(r.Icon)}
		} else {
			publication["icon"] = nil
		}

		if r.Id != nil {
			publication["id"] = *r.Id
		}

		if r.Name != nil {
			publication["name"] = *r.Name
		}

		publication["package_type"] = r.PackageType

		if r.ShortDescription != nil {
			publication["short_description"] = *r.ShortDescription
		}

		publication["state"] = r.LifecycleState

		supportedOperatingSystems := []interface{}{}
		for _, item := range r.SupportedOperatingSystems {
			supportedOperatingSystems = append(supportedOperatingSystems, OperatingSystemToMap(item))
		}
		publication["supported_operating_systems"] = supportedOperatingSystems

		if r.TimeCreated != nil {
			publication["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, publication)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MarketplacePublicationsDataSource().Schema["publications"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("publications", resources); err != nil {
		return err
	}

	return nil
}
