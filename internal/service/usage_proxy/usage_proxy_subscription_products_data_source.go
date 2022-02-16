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

func UsageProxySubscriptionProductsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsageProxySubscriptionProducts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"producttype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usage_period_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"product_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"earned_rewards": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"is_eligible_to_earn_rewards": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"product_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"product_number": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"usage_amount": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readUsageProxySubscriptionProducts(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionProductsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxySubscriptionProductsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListProductsResponse
}

func (s *UsageProxySubscriptionProductsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionProductsDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListProductsRequest{}

	if producttype, ok := s.D.GetOkExists("producttype"); ok {
		request.Producttype = oci_usage_proxy.ListProductsProducttypeEnum(producttype.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	if usagePeriodKey, ok := s.D.GetOkExists("usage_period_key"); ok {
		tmp := usagePeriodKey.(string)
		request.UsagePeriodKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListProducts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProducts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UsageProxySubscriptionProductsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxySubscriptionProductsDataSource-", UsageProxySubscriptionProductsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subscriptionProduct := map[string]interface{}{}

	itemmap := map[string]interface{}{}
	items := []interface{}{}
	resources2 := []map[string]interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProductSummaryToMap(item))
	}
	subscriptionProduct["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {

		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, UsageProxySubscriptionProductsDataSource().Schema["product_collection"].Elem.(*schema.Resource).Schema)
		subscriptionProduct["items"] = items
	}

	resources = append(resources, subscriptionProduct)
	itemmap["items"] = resources

	resources2 = append(resources2, itemmap)

	if err := s.D.Set("product_collection", resources2); err != nil {
		return err
	}

	return nil
}

func ProductSummaryToMap(obj oci_usage_proxy.ProductSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EarnedRewards != nil {
		result["earned_rewards"] = float32(*obj.EarnedRewards)
	}

	if obj.IsEligibleToEarnRewards != nil {
		result["is_eligible_to_earn_rewards"] = bool(*obj.IsEligibleToEarnRewards)
	}

	if obj.ProductName != nil {
		result["product_name"] = string(*obj.ProductName)
	}

	if obj.ProductNumber != nil {
		result["product_number"] = string(*obj.ProductNumber)
	}

	if obj.UsageAmount != nil {
		result["usage_amount"] = float64(*obj.UsageAmount)
	}

	return result
}
