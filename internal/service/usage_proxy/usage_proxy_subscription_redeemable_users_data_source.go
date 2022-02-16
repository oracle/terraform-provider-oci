// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package usage_proxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v58/usage"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func UsageProxySubscriptionRedeemableUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsageProxySubscriptionRedeemableUsers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"redeemable_user_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     UsageProxySubscriptionRedeemableUserResource(),
						},
					},
				},
			},
		},
	}
}

func readUsageProxySubscriptionRedeemableUsers(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRedeemableUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxySubscriptionRedeemableUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListRedeemableUsersResponse
}

func (s *UsageProxySubscriptionRedeemableUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionRedeemableUsersDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListRedeemableUsersRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListRedeemableUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRedeemableUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UsageProxySubscriptionRedeemableUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxySubscriptionRedeemableUsersDataSource-", UsageProxySubscriptionRedeemableUsersDataSource(), s.D))
	resources := []map[string]interface{}{}
	subscriptionRedeemableUser := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RedeemableUserSummaryToMap(item))
	}
	subscriptionRedeemableUser["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, UsageProxySubscriptionRedeemableUsersDataSource().Schema["redeemable_user_collection"].Elem.(*schema.Resource).Schema)
		subscriptionRedeemableUser["items"] = items
	}

	resources = append(resources, subscriptionRedeemableUser)
	if err := s.D.Set("redeemable_user_collection", resources); err != nil {
		return err
	}

	return nil
}
