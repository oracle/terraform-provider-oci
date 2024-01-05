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

func OsubUsageComputedUsagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsubUsageComputedUsages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"computed_product": {
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
			"computed_usages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"computed_usage_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"commitment_service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compute_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cost": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cost_rounded": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"currency_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_center": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_invoiced": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mqs_message_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"net_unit_price": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"original_usage_number": {
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
						"rate_card_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rate_card_tierd_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_metered_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_of_arrival": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"unit_of_measure": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"usage_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOsubUsageComputedUsages(d *schema.ResourceData, m interface{}) error {
	sync := &OsubUsageComputedUsagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputedUsageClient()

	return tfresource.ReadResource(sync)
}

type OsubUsageComputedUsagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_usage.ComputedUsageClient
	Res    *oci_osub_usage.ListComputedUsagesResponse
}

func (s *OsubUsageComputedUsagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubUsageComputedUsagesDataSourceCrud) Get() error {
	request := oci_osub_usage.ListComputedUsagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computedProduct, ok := s.D.GetOkExists("computed_product"); ok {
		tmp := computedProduct.(string)
		request.ComputedProduct = &tmp
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

	response, err := s.Client.ListComputedUsages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputedUsages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsubUsageComputedUsagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsubUsageComputedUsagesDataSource-", OsubUsageComputedUsagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computedUsage := map[string]interface{}{}

		if r.CommitmentServiceId != nil {
			computedUsage["commitment_service_id"] = *r.CommitmentServiceId
		}

		if r.ComputeSource != nil {
			computedUsage["compute_source"] = *r.ComputeSource
		}

		if r.ComputedUsageId != nil {
			computedUsage["computed_usage_id"] = *r.ComputedUsageId
		}

		if r.Cost != nil {
			computedUsage["cost"] = *r.Cost
		}

		if r.CostRounded != nil {
			computedUsage["cost_rounded"] = *r.CostRounded
		}

		if r.CurrencyCode != nil {
			computedUsage["currency_code"] = *r.CurrencyCode
		}

		if r.DataCenter != nil {
			computedUsage["data_center"] = *r.DataCenter
		}

		if r.IsInvoiced != nil {
			computedUsage["is_invoiced"] = *r.IsInvoiced
		}

		if r.MqsMessageId != nil {
			computedUsage["mqs_message_id"] = *r.MqsMessageId
		}

		if r.NetUnitPrice != nil {
			computedUsage["net_unit_price"] = *r.NetUnitPrice
		}

		if r.OriginalUsageNumber != nil {
			computedUsage["original_usage_number"] = *r.OriginalUsageNumber
		}

		if r.ParentProduct != nil {
			computedUsage["parent_product"] = []interface{}{ProductToMap(r.ParentProduct)}
		} else {
			computedUsage["parent_product"] = nil
		}

		if r.ParentSubscribedServiceId != nil {
			computedUsage["parent_subscribed_service_id"] = *r.ParentSubscribedServiceId
		}

		if r.PlanNumber != nil {
			computedUsage["plan_number"] = *r.PlanNumber
		}

		if r.Product != nil {
			computedUsage["product"] = []interface{}{ProductToMap(r.Product)}
		} else {
			computedUsage["product"] = nil
		}

		if r.Quantity != nil {
			computedUsage["quantity"] = *r.Quantity
		}

		if r.RateCardId != nil {
			computedUsage["rate_card_id"] = *r.RateCardId
		}

		if r.RateCardTierdId != nil {
			computedUsage["rate_card_tierd_id"] = *r.RateCardTierdId
		}

		if r.TimeCreated != nil {
			computedUsage["time_created"] = r.TimeCreated.String()
		}

		if r.TimeMeteredOn != nil {
			computedUsage["time_metered_on"] = r.TimeMeteredOn.String()
		}

		if r.TimeOfArrival != nil {
			computedUsage["time_of_arrival"] = r.TimeOfArrival.String()
		}

		if r.TimeUpdated != nil {
			computedUsage["time_updated"] = r.TimeUpdated.String()
		}

		computedUsage["type"] = r.Type

		if r.UnitOfMeasure != nil {
			computedUsage["unit_of_measure"] = *r.UnitOfMeasure
		}

		if r.UsageNumber != nil {
			computedUsage["usage_number"] = *r.UsageNumber
		}

		resources = append(resources, computedUsage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsubUsageComputedUsagesDataSource().Schema["computed_usages"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("computed_usages", resources); err != nil {
		return err
	}

	return nil
}

func ProductToMap(obj *oci_osub_usage.Product) map[string]interface{} {
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

	if obj.ProvisioningGroup != nil {
		result["provisioning_group"] = string(*obj.ProvisioningGroup)
	}

	if obj.UcmRateCardPartType != nil {
		result["ucm_rate_card_part_type"] = string(*obj.UcmRateCardPartType)
	}

	if obj.UnitOfMeasure != nil {
		result["unit_of_measure"] = string(*obj.UnitOfMeasure)
	}

	return result
}
