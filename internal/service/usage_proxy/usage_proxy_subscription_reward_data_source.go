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

func UsageProxySubscriptionRewardDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularUsageProxySubscriptionReward,
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenancy_id": {
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
						"available_rewards": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"earned_rewards": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"eligible_usage_amount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"ineligible_usage_amount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"is_manual": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"redeemed_rewards": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"time_rewards_earned": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_rewards_expired": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_usage_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_usage_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"usage_amount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"usage_period_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"summary": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"currency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"redemption_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rewards_rate": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"subscription_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenancy_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_rewards_available": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_usage_proxy_subscription_reward", "oci_usage_proxy_subscription_rewards"),
	}
}

func readSingularUsageProxySubscriptionReward(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRewardDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxySubscriptionRewardDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListRewardsResponse
}

func (s *UsageProxySubscriptionRewardDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionRewardDataSourceCrud) Get() error {
	request := oci_usage_proxy.ListRewardsRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "usage_proxy")

	response, err := s.Client.ListRewards(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *UsageProxySubscriptionRewardDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxySubscriptionRewardDataSource-", UsageProxySubscriptionRewardDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonthlyRewardSummaryToMap(item))
	}
	s.D.Set("items", items)

	if s.Res.Summary != nil {
		s.D.Set("summary", []interface{}{RewardDetailsToMap(s.Res.Summary)})
	} else {
		s.D.Set("summary", nil)
	}

	return nil
}
