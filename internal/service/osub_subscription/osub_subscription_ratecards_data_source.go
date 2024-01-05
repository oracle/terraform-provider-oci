// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_subscription

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_osub_subscription "github.com/oracle/oci-go-sdk/v65/osubsubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsubSubscriptionRatecardsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsubSubscriptionRatecards,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_from": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"x_one_origin_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rate_cards": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"currency": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"iso_code": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"std_precision": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"discretionary_discount_percentage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_tier": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"net_unit_price": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overage_price": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"product": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"billing_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"part_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ucm_rate_card_part_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"unit_of_measure": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"rate_card_tiers": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"net_unit_price": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"overage_price": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"up_to_quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOsubSubscriptionRatecards(d *schema.ResourceData, m interface{}) error {
	sync := &OsubSubscriptionRatecardsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RatecardClient()

	return tfresource.ReadResource(sync)
}

type OsubSubscriptionRatecardsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_subscription.RatecardClient
	Res    *oci_osub_subscription.ListRateCardsResponse
}

func (s *OsubSubscriptionRatecardsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubSubscriptionRatecardsDataSourceCrud) Get() error {
	request := oci_osub_subscription.ListRateCardsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if partNumber, ok := s.D.GetOkExists("part_number"); ok {
		tmp := partNumber.(string)
		request.PartNumber = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if timeFrom, ok := s.D.GetOkExists("time_from"); ok {
		tmp, err := time.Parse(time.RFC3339, timeFrom.(string))
		if err != nil {
			return err
		}
		request.TimeFrom = &oci_common.SDKTime{Time: tmp}
	}

	if timeTo, ok := s.D.GetOkExists("time_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeTo.(string))
		if err != nil {
			return err
		}
		request.TimeTo = &oci_common.SDKTime{Time: tmp}
	}

	if xOneOriginRegion, ok := s.D.GetOkExists("x_one_origin_region"); ok {
		tmp := xOneOriginRegion.(string)
		request.XOneOriginRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osub_subscription")

	response, err := s.Client.ListRateCards(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRateCards(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsubSubscriptionRatecardsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsubSubscriptionRatecardsDataSource-", OsubSubscriptionRatecardsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ratecard := map[string]interface{}{}

		if r.Currency != nil {
			ratecard["currency"] = []interface{}{CurrencyToMap(r.Currency)}
		} else {
			ratecard["currency"] = nil
		}

		if r.DiscretionaryDiscountPercentage != nil {
			ratecard["discretionary_discount_percentage"] = *r.DiscretionaryDiscountPercentage
		}

		if r.IsTier != nil {
			ratecard["is_tier"] = *r.IsTier
		}

		if r.NetUnitPrice != nil {
			ratecard["net_unit_price"] = *r.NetUnitPrice
		}

		if r.OveragePrice != nil {
			ratecard["overage_price"] = *r.OveragePrice
		}

		if r.Product != nil {
			ratecard["product"] = []interface{}{ProductToMap(r.Product)}
		} else {
			ratecard["product"] = nil
		}

		rateCardTiers := []interface{}{}
		for _, item := range r.RateCardTiers {
			rateCardTiers = append(rateCardTiers, RateCardTierToMap(item))
		}
		ratecard["rate_card_tiers"] = rateCardTiers

		if r.TimeEnd != nil {
			ratecard["time_end"] = r.TimeEnd.String()
		}

		if r.TimeStart != nil {
			ratecard["time_start"] = r.TimeStart.String()
		}

		resources = append(resources, ratecard)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsubSubscriptionRatecardsDataSource().Schema["rate_cards"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("rate_cards", resources); err != nil {
		return err
	}

	return nil
}

func CurrencyToMap(obj *oci_osub_subscription.Currency) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsoCode != nil {
		result["iso_code"] = string(*obj.IsoCode)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.StdPrecision != nil {
		result["std_precision"] = strconv.FormatInt(*obj.StdPrecision, 10)
	}

	return result
}

func ProductToMap(obj *oci_osub_subscription.Product) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillingCategory != nil {
		result["billing_category"] = string(*obj.BillingCategory)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PartNumber != nil {
		result["part_number"] = string(*obj.PartNumber)
	}

	if obj.ProductCategory != nil {
		result["product_category"] = string(*obj.ProductCategory)
	}

	if obj.UcmRateCardPartType != nil {
		result["ucm_rate_card_part_type"] = string(*obj.UcmRateCardPartType)
	}

	if obj.UnitOfMeasure != nil {
		result["unit_of_measure"] = string(*obj.UnitOfMeasure)
	}

	return result
}

func RateCardTierToMap(obj oci_osub_subscription.RateCardTier) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NetUnitPrice != nil {
		result["net_unit_price"] = string(*obj.NetUnitPrice)
	}

	if obj.OveragePrice != nil {
		result["overage_price"] = string(*obj.OveragePrice)
	}

	if obj.UpToQuantity != nil {
		result["up_to_quantity"] = string(*obj.UpToQuantity)
	}

	return result
}
