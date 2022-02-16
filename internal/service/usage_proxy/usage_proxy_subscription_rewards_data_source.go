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

func UsageProxySubscriptionRewardsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsageProxySubscriptionRewards,
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
			"reward_collection": {
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
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"currency": {
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
							},
						},
					},
				},
			},
		},
	}
}

func readUsageProxySubscriptionRewards(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRewardsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxySubscriptionRewardsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListRewardsResponse
}

func (s *UsageProxySubscriptionRewardsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionRewardsDataSourceCrud) Get() error {
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

func (s *UsageProxySubscriptionRewardsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxySubscriptionRewardsDataSource-", UsageProxySubscriptionRewardsDataSource(), s.D))
	resources := []map[string]interface{}{}
	resources2 := []map[string]interface{}{}
	subscriptionReward := map[string]interface{}{}
	itemmap := map[string]interface{}{}

	items := []interface{}{}

	for _, item := range s.Res.Items {
		items = append(items, MonthlyRewardSummaryToMap(item))
	}
	subscriptionReward["items"] = items

	if s.Res.Summary != nil {
		subscriptionReward["summary"] = []interface{}{RewardDetailsToMap(s.Res.Summary)}
	} else {
		subscriptionReward["summary"] = nil
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, UsageProxySubscriptionRewardsDataSource().Schema["reward_collection"].Elem.(*schema.Resource).Schema)
		subscriptionReward["items"] = items
	}

	resources = append(resources, subscriptionReward)

	itemmap["items"] = resources

	resources2 = append(resources2, itemmap)

	if err := s.D.Set("reward_collection", resources2); err != nil {
		return err
	}

	return nil
}

func MonthlyRewardSummaryToMap(obj oci_usage_proxy.MonthlyRewardSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableRewards != nil {
		result["available_rewards"] = float32(*obj.AvailableRewards)
	}

	if obj.EarnedRewards != nil {
		result["earned_rewards"] = float32(*obj.EarnedRewards)
	}

	if obj.EligibleUsageAmount != nil {
		result["eligible_usage_amount"] = float64(*obj.EligibleUsageAmount)
	}

	if obj.IneligibleUsageAmount != nil {
		result["ineligible_usage_amount"] = float64(*obj.IneligibleUsageAmount)
	}

	if obj.IsManual != nil {
		result["is_manual"] = bool(*obj.IsManual)
	}

	if obj.RedeemedRewards != nil {
		result["redeemed_rewards"] = float32(*obj.RedeemedRewards)
	}

	if obj.TimeRewardsEarned != nil {
		result["time_rewards_earned"] = obj.TimeRewardsEarned.String()
	}

	if obj.TimeRewardsExpired != nil {
		result["time_rewards_expired"] = obj.TimeRewardsExpired.String()
	}

	if obj.TimeUsageEnded != nil {
		result["time_usage_ended"] = obj.TimeUsageEnded.String()
	}

	if obj.TimeUsageStarted != nil {
		result["time_usage_started"] = obj.TimeUsageStarted.String()
	}

	if obj.UsageAmount != nil {
		result["usage_amount"] = float64(*obj.UsageAmount)
	}

	if obj.UsagePeriodKey != nil {
		result["usage_period_key"] = string(*obj.UsagePeriodKey)
	}

	return result
}

func RewardDetailsToMap(obj *oci_usage_proxy.RewardDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Currency != nil {
		result["currency"] = string(*obj.Currency)
	}

	if obj.RewardsRate != nil {
		result["rewards_rate"] = float64(*obj.RewardsRate)
	}

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.TotalRewardsAvailable != nil {
		result["total_rewards_available"] = float32(*obj.TotalRewardsAvailable)
	}

	return result
}
