// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreAppCatalogSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreAppCatalogSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_catalog_subscriptions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CoreAppCatalogSubscriptionResource(),
			},
		},
	}
}

func readCoreAppCatalogSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreAppCatalogSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreAppCatalogSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListAppCatalogSubscriptionsResponse
}

func (s *CoreAppCatalogSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreAppCatalogSubscriptionsDataSourceCrud) Get() error {
	request := oci_core.ListAppCatalogSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListAppCatalogSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAppCatalogSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreAppCatalogSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreAppCatalogSubscriptionsDataSource-", CoreAppCatalogSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		appCatalogSubscription := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			appCatalogSubscription["display_name"] = *r.DisplayName
		}

		if r.ListingId != nil {
			appCatalogSubscription["listing_id"] = *r.ListingId
		}

		if r.ListingResourceId != nil {
			appCatalogSubscription["listing_resource_id"] = *r.ListingResourceId
		}

		if r.ListingResourceVersion != nil {
			appCatalogSubscription["listing_resource_version"] = *r.ListingResourceVersion
		}

		if r.PublisherName != nil {
			appCatalogSubscription["publisher_name"] = *r.PublisherName
		}

		if r.Summary != nil {
			appCatalogSubscription["summary"] = *r.Summary
		}

		if r.TimeCreated != nil {
			appCatalogSubscription["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, appCatalogSubscription)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreAppCatalogSubscriptionsDataSource().Schema["app_catalog_subscriptions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("app_catalog_subscriptions", resources); err != nil {
		return err
	}

	return nil
}
