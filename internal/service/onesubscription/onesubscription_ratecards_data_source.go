// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OnesubscriptionRatecardsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionRatecards,
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
						"subscribed_service_id": {
							Type:     schema.TypeString,
							Computed: true,
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

func readOnesubscriptionRatecards(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionRatecardsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).RatecardRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionRatecardsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.RatecardClient
	Res    *oci_onesubscription.ListRateCardsResponse
}

func (s *OnesubscriptionRatecardsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionRatecardsDataSourceCrud) Get() error {
	request := oci_onesubscription.ListRateCardsRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

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

func (s *OnesubscriptionRatecardsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionRatecardsDataSource-", OnesubscriptionRatecardsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		ratecard := map[string]interface{}{}

		if r.Currency != nil {
			ratecard["currency"] = []interface{}{SubscriptionCurrencyToMap(r.Currency)}
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
			ratecard["product"] = []interface{}{RateCardProductToMap(r.Product)}
		} else {
			ratecard["product"] = nil
		}

		rateCardTiers := []interface{}{}
		for _, item := range r.RateCardTiers {
			rateCardTiers = append(rateCardTiers, RateCardTierToMap(item))
		}
		ratecard["rate_card_tiers"] = rateCardTiers

		if r.SubscribedServiceId != nil {
			ratecard["subscribed_service_id"] = *r.SubscribedServiceId
		}

		if r.TimeEnd != nil {
			ratecard["time_end"] = r.TimeEnd.String()
		}

		if r.TimeStart != nil {
			ratecard["time_start"] = r.TimeStart.String()
		}

		resources = append(resources, ratecard)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionRatecardsDataSource().Schema["rate_cards"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("rate_cards", resources); err != nil {
		return err
	}

	return nil
}
