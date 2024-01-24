// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package usage_proxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v65/usage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func UsageProxySubscriptionProductDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularUsageProxySubscriptionProduct,
		Schema: map[string]*schema.Schema{
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_usage_proxy_subscription_product", "oci_usage_proxy_subscription_products"),
	}
}

func readSingularUsageProxySubscriptionProduct(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionProductDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxySubscriptionProductDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListProductsResponse
}

func (s *UsageProxySubscriptionProductDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionProductDataSourceCrud) Get() error {
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
	return nil
}

func (s *UsageProxySubscriptionProductDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxySubscriptionProductDataSource-", UsageProxySubscriptionProductDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProductSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
