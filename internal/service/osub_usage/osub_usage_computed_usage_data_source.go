// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_usage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osub_usage "github.com/oracle/oci-go-sdk/v65/osubusage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsubUsageComputedUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsubUsageComputedUsage,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"computed_usage_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"x_one_origin_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
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
	}
}

func readSingularOsubUsageComputedUsage(d *schema.ResourceData, m interface{}) error {
	sync := &OsubUsageComputedUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputedUsageClient()

	return tfresource.ReadResource(sync)
}

type OsubUsageComputedUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_usage.ComputedUsageClient
	Res    *oci_osub_usage.GetComputedUsageResponse
}

func (s *OsubUsageComputedUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubUsageComputedUsageDataSourceCrud) Get() error {
	request := oci_osub_usage.GetComputedUsageRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computedUsageId, ok := s.D.GetOkExists("computed_usage_id"); ok {
		tmp := computedUsageId.(string)
		request.ComputedUsageId = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if xOneOriginRegion, ok := s.D.GetOkExists("x_one_origin_region"); ok {
		tmp := xOneOriginRegion.(string)
		request.XOneOriginRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osub_usage")

	response, err := s.Client.GetComputedUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsubUsageComputedUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CommitmentServiceId != nil {
		s.D.Set("commitment_service_id", *s.Res.CommitmentServiceId)
	}

	if s.Res.ComputeSource != nil {
		s.D.Set("compute_source", *s.Res.ComputeSource)
	}

	if s.Res.Cost != nil {
		s.D.Set("cost", *s.Res.Cost)
	}

	if s.Res.CostRounded != nil {
		s.D.Set("cost_rounded", *s.Res.CostRounded)
	}

	if s.Res.CurrencyCode != nil {
		s.D.Set("currency_code", *s.Res.CurrencyCode)
	}

	if s.Res.DataCenter != nil {
		s.D.Set("data_center", *s.Res.DataCenter)
	}

	if s.Res.IsInvoiced != nil {
		s.D.Set("is_invoiced", *s.Res.IsInvoiced)
	}

	if s.Res.MqsMessageId != nil {
		s.D.Set("mqs_message_id", *s.Res.MqsMessageId)
	}

	if s.Res.NetUnitPrice != nil {
		s.D.Set("net_unit_price", *s.Res.NetUnitPrice)
	}

	if s.Res.OriginalUsageNumber != nil {
		s.D.Set("original_usage_number", *s.Res.OriginalUsageNumber)
	}

	if s.Res.ParentProduct != nil {
		s.D.Set("parent_product", []interface{}{ProductToMap(s.Res.ParentProduct)})
	} else {
		s.D.Set("parent_product", nil)
	}

	if s.Res.ParentSubscribedServiceId != nil {
		s.D.Set("parent_subscribed_service_id", *s.Res.ParentSubscribedServiceId)
	}

	if s.Res.PlanNumber != nil {
		s.D.Set("plan_number", *s.Res.PlanNumber)
	}

	if s.Res.Product != nil {
		s.D.Set("product", []interface{}{ProductToMap(s.Res.Product)})
	} else {
		s.D.Set("product", nil)
	}

	if s.Res.Quantity != nil {
		s.D.Set("quantity", *s.Res.Quantity)
	}

	if s.Res.RateCardId != nil {
		s.D.Set("rate_card_id", *s.Res.RateCardId)
	}

	if s.Res.RateCardTierdId != nil {
		s.D.Set("rate_card_tierd_id", *s.Res.RateCardTierdId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeMeteredOn != nil {
		s.D.Set("time_metered_on", s.Res.TimeMeteredOn.String())
	}

	if s.Res.TimeOfArrival != nil {
		s.D.Set("time_of_arrival", s.Res.TimeOfArrival.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UnitOfMeasure != nil {
		s.D.Set("unit_of_measure", *s.Res.UnitOfMeasure)
	}

	if s.Res.UsageNumber != nil {
		s.D.Set("usage_number", *s.Res.UsageNumber)
	}

	return nil
}
