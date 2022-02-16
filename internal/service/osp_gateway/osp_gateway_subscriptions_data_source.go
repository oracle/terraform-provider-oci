// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v58/ospgateway"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OspGatewaySubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOspGatewaySubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OspGatewaySubscriptionResource()),
						},
					},
				},
			},
		},
	}
}

func readOspGatewaySubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewaySubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscriptionServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewaySubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.SubscriptionServiceClient
	Res    *oci_osp_gateway.ListSubscriptionsResponse
}

func (s *OspGatewaySubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewaySubscriptionsDataSourceCrud) Get() error {
	request := oci_osp_gateway.ListSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osp_gateway")

	response, err := s.Client.ListSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OspGatewaySubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OspGatewaySubscriptionsDataSource-", OspGatewaySubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subscription := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubscriptionSummaryToMap(item))
	}
	subscription["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OspGatewaySubscriptionsDataSource().Schema["subscription_collection"].Elem.(*schema.Resource).Schema)
		subscription["items"] = items
	}

	resources = append(resources, subscription)
	if err := s.D.Set("subscription_collection", resources); err != nil {
		return err
	}

	return nil
}
