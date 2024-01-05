// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_usage

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_osub_usage "github.com/oracle/oci-go-sdk/v65/osubusage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsubUsageComputedUsageAggregatedsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsubUsageComputedUsageAggregateds,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"grouping": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_product": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_from": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_to": {
				Type:     schema.TypeString,
				Required: true,
			},
			"x_one_origin_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"computed_usage_aggregateds": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"aggregated_computed_usages": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"cost": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cost_unrounded": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data_center": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"net_unit_price": {
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
												"provisioning_group": {
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
									"quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_metered_on": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"currency_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_product": {
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
									"provisioning_group": {
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
						"parent_subscribed_service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plan_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pricing_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rate_card_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_id": {
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

func readOsubUsageComputedUsageAggregateds(d *schema.ResourceData, m interface{}) error {
	sync := &OsubUsageComputedUsageAggregatedsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputedUsageClient()

	return tfresource.ReadResource(sync)
}

type OsubUsageComputedUsageAggregatedsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_usage.ComputedUsageClient
	Res    *oci_osub_usage.ListComputedUsageAggregatedsResponse
}

func (s *OsubUsageComputedUsageAggregatedsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubUsageComputedUsageAggregatedsDataSourceCrud) Get() error {
	request := oci_osub_usage.ListComputedUsageAggregatedsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if grouping, ok := s.D.GetOkExists("grouping"); ok {
		request.Grouping = oci_osub_usage.ListComputedUsageAggregatedsGroupingEnum(grouping.(string))
	}

	if parentProduct, ok := s.D.GetOkExists("parent_product"); ok {
		tmp := parentProduct.(string)
		request.ParentProduct = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osub_usage")

	response, err := s.Client.ListComputedUsageAggregateds(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputedUsageAggregateds(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsubUsageComputedUsageAggregatedsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsubUsageComputedUsageAggregatedsDataSource-", OsubUsageComputedUsageAggregatedsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computedUsageAggregated := map[string]interface{}{
			"subscription_id": *r.SubscriptionId,
		}

		aggregatedComputedUsages := []interface{}{}
		for _, item := range r.AggregatedComputedUsages {
			aggregatedComputedUsages = append(aggregatedComputedUsages, ComputedUsageAggregationToMap(item))
		}
		computedUsageAggregated["aggregated_computed_usages"] = aggregatedComputedUsages

		if r.CurrencyCode != nil {
			computedUsageAggregated["currency_code"] = *r.CurrencyCode
		}

		if r.ParentProduct != nil {
			computedUsageAggregated["parent_product"] = []interface{}{ProductToMap(r.ParentProduct)}
		} else {
			computedUsageAggregated["parent_product"] = nil
		}

		if r.ParentSubscribedServiceId != nil {
			computedUsageAggregated["parent_subscribed_service_id"] = *r.ParentSubscribedServiceId
		}

		if r.PlanNumber != nil {
			computedUsageAggregated["plan_number"] = *r.PlanNumber
		}

		computedUsageAggregated["pricing_model"] = r.PricingModel

		if r.RateCardId != nil {
			computedUsageAggregated["rate_card_id"] = *r.RateCardId
		}

		if r.TimeEnd != nil {
			computedUsageAggregated["time_end"] = r.TimeEnd.String()
		}

		if r.TimeStart != nil {
			computedUsageAggregated["time_start"] = r.TimeStart.String()
		}

		resources = append(resources, computedUsageAggregated)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsubUsageComputedUsageAggregatedsDataSource().Schema["computed_usage_aggregateds"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("computed_usage_aggregateds", resources); err != nil {
		return err
	}

	return nil
}

func ComputedUsageAggregationToMap(obj oci_osub_usage.ComputedUsageAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cost != nil {
		result["cost"] = string(*obj.Cost)
	}

	if obj.CostUnrounded != nil {
		result["cost_unrounded"] = string(*obj.CostUnrounded)
	}

	if obj.DataCenter != nil {
		result["data_center"] = string(*obj.DataCenter)
	}

	if obj.NetUnitPrice != nil {
		result["net_unit_price"] = string(*obj.NetUnitPrice)
	}

	if obj.Product != nil {
		result["product"] = []interface{}{ProductToMap(obj.Product)}
	}

	if obj.Quantity != nil {
		result["quantity"] = string(*obj.Quantity)
	}

	if obj.TimeMeteredOn != nil {
		result["time_metered_on"] = obj.TimeMeteredOn.String()
	}

	result["type"] = string(obj.Type)

	return result
}
