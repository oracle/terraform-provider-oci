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

func MarketplacePublishersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplacePublishers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publisher_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"publishers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMarketplacePublishers(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplacePublishersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplacePublishersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListPublishersResponse
}

func (s *MarketplacePublishersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplacePublishersDataSourceCrud) Get() error {
	request := oci_marketplace.ListPublishersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if publisherId, ok := s.D.GetOkExists("publisher_id"); ok {
		tmp := publisherId.(string)
		request.PublisherId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.ListPublishers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPublishers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplacePublishersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MarketplacePublishersDataSource-", MarketplacePublishersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		publisher := map[string]interface{}{}

		if r.Description != nil {
			publisher["description"] = *r.Description
		}

		if r.Id != nil {
			publisher["id"] = *r.Id
		}

		if r.Name != nil {
			publisher["name"] = *r.Name
		}

		resources = append(resources, publisher)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MarketplacePublishersDataSource().Schema["publishers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("publishers", resources); err != nil {
		return err
	}

	return nil
}
