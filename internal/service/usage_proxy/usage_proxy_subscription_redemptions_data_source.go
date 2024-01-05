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

func UsageProxySubscriptionRedemptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readUsageProxySubscriptionRedemptions,
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
			"time_redeemed_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_redeemed_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redemption_collection": {
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
							},
						},
					},
				},
			},
		},
	}
}

func readUsageProxySubscriptionRedemptions(d *schema.ResourceData, m interface{}) error {
	sync := &UsageProxySubscriptionRedemptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RewardsClient()

	return tfresource.ReadResource(sync)
}

type UsageProxySubscriptionRedemptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_usage_proxy.RewardsClient
	Res    *oci_usage_proxy.ListRedemptionsResponse
}

func (s *UsageProxySubscriptionRedemptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *UsageProxySubscriptionRedemptionsDataSourceCrud) Get() error {
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
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRedemptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *UsageProxySubscriptionRedemptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("UsageProxySubscriptionRedemptionsDataSource-", UsageProxySubscriptionRedemptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	resources2 := []map[string]interface{}{}
	subscriptionRedemption := map[string]interface{}{}
	itemmap := map[string]interface{}{}

	items := []interface{}{}

	for _, item := range s.Res.Items {
		items = append(items, RedemptionSummaryToMap(item))
	}
	subscriptionRedemption["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, UsageProxySubscriptionRedemptionsDataSource().Schema["redemption_collection"].Elem.(*schema.Resource).Schema)
		subscriptionRedemption["items"] = items
	}

	resources = append(resources, subscriptionRedemption)
	itemmap["items"] = resources
	resources2 = append(resources2, itemmap)
	if err := s.D.Set("redemption_collection", resources2); err != nil {
		return err
	}

	return nil
}

func RedemptionSummaryToMap(obj oci_usage_proxy.RedemptionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRewards != nil {
		result["base_rewards"] = float32(*obj.BaseRewards)
	}

	if obj.FxRate != nil {
		result["fx_rate"] = float64(*obj.FxRate)
	}

	if obj.InvoiceCurrency != nil {
		result["invoice_currency"] = string(*obj.InvoiceCurrency)
	}

	if obj.InvoiceNumber != nil {
		result["invoice_number"] = string(*obj.InvoiceNumber)
	}

	if obj.InvoiceTotalAmount != nil {
		result["invoice_total_amount"] = float64(*obj.InvoiceTotalAmount)
	}

	if obj.RedeemedRewards != nil {
		result["redeemed_rewards"] = float32(*obj.RedeemedRewards)
	}

	if obj.RedemptionCode != nil {
		result["redemption_code"] = string(*obj.RedemptionCode)
	}

	if obj.RedemptionEmail != nil {
		result["redemption_email"] = string(*obj.RedemptionEmail)
	}

	if obj.TimeInvoiced != nil {
		result["time_invoiced"] = obj.TimeInvoiced.String()
	}

	if obj.TimeRedeemed != nil {
		result["time_redeemed"] = obj.TimeRedeemed.String()
	}

	return result
}
