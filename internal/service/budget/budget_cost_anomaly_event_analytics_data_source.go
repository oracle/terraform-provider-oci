// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BudgetCostAnomalyEventAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readBudgetCostAnomalyEventAnalyticsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cost_anomaly_monitor_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost_impact": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"cost_impact_percentage": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_tenant_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_anomaly_event_end_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_anomaly_event_start_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost_anomaly_event_analytic_collection": {
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
									"average_cost_impact": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"average_cost_variance": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"cost_anomaly_event_analytic_count": {
										Type:     schema.TypeInt,
										Computed: true,
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

func readBudgetCostAnomalyEventAnalyticsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAnomalyEventAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BudgetCostAnomalyEventAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_budget.CostAdClient
	Res    *oci_budget.SummarizeCostAnomalyEventAnalyticsResponse
}

func (s *BudgetCostAnomalyEventAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BudgetCostAnomalyEventAnalyticsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_budget.SummarizeCostAnomalyEventAnalyticsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	if costImpact, ok := s.D.GetOkExists("cost_impact"); ok {
		tmp := costImpact.(float64)
		request.CostImpact = &tmp
	}

	//if costImpactPercentage, ok := s.D.GetOkExists("cost_impact_percentage"); ok {
	//	tmp := costImpactPercentage.(float64)
	//	request.CostImpactPercentage = &tmp
	//}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		interfaces := region.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("region") {
			request.Region = tmp
		}
	}

	if targetTenantId, ok := s.D.GetOkExists("target_tenant_id"); ok {
		interfaces := targetTenantId.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("target_tenant_id") {
			request.TargetTenantId = tmp
		}
	}

	if timeAnomalyEventEndDate, ok := s.D.GetOkExists("time_anomaly_event_end_date"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAnomalyEventEndDate.(string))
		if err != nil {
			return err
		}
		request.TimeAnomalyEventEndDate = &oci_common.SDKTime{Time: tmp}
	}

	if timeAnomalyEventStartDate, ok := s.D.GetOkExists("time_anomaly_event_start_date"); ok {
		tmp, err := time.Parse(time.RFC3339, timeAnomalyEventStartDate.(string))
		if err != nil {
			return err
		}
		request.TimeAnomalyEventStartDate = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "budget")

	response, err := s.Client.SummarizeCostAnomalyEventAnalytics(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.SummarizeCostAnomalyEventAnalytics(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BudgetCostAnomalyEventAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BudgetCostAnomalyEventAnalyticsDataSource-", BudgetCostAnomalyEventAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	costAnomalyEventAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CostAnomalyEventAnalyticSummaryToMap(item))
	}
	costAnomalyEventAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BudgetCostAnomalyEventAnalyticsDataSource().Schema["cost_anomaly_event_analytic_collection"].Elem.(*schema.Resource).Schema)
		costAnomalyEventAnalytic["items"] = items
	}

	resources = append(resources, costAnomalyEventAnalytic)
	if err := s.D.Set("cost_anomaly_event_analytic_collection", resources); err != nil {
		return err
	}

	return nil
}

func CostAnomalyEventAnalyticSummaryToMap(obj oci_budget.CostAnomalyEventAnalyticSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AverageCostImpact != nil {
		result["average_cost_impact"] = float64(*obj.AverageCostImpact)
	}

	if obj.AverageCostVariance != nil {
		result["average_cost_variance"] = float64(*obj.AverageCostVariance)
	}

	if obj.Count != nil {
		result["cost_anomaly_event_analytic_count"] = int(*obj.Count)
	}

	return result
}
