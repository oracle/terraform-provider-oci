// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package usage_proxy

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v65/usage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func UsageProxySubscriptionRedemptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularUsageProxySubscriptionRedemption,
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_redeemed_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_redeemed_less_than": {
				Type:     schema.TypeString,
				Optional: true,
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
						"base_rewards": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"fx_rate": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"invoice_currency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"invoice_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"invoice_total_amount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"redeemed_rewards": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"redemption_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"redemption_email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_invoiced": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_redeemed": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_usage_proxy_subscription_redemption", "oci_usage_proxy_subscription_redemptions"),
	}
}

func readSingularUsageProxySubscriptionRedemption(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRedemptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxySubscriptionRedemptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListRedemptionsResponse
}

func (s *UsageProxySubscriptionRedemptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionRedemptionDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListRedemptionsRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	if timeRedeemedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_redeemed_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeRedeemedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeRedeemedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeRedeemedLessThan, ok := s.D.GetOkExists("time_redeemed_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeRedeemedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeRedeemedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListRedemptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *UsageProxySubscriptionRedemptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxySubscriptionRedemptionDataSource-", UsageProxySubscriptionRedemptionDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RedemptionSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
