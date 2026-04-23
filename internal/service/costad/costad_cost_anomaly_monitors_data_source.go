// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package costad

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_costad "github.com/oracle/oci-go-sdk/v65/costad"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CostadCostAnomalyMonitorsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readCostadCostAnomalyMonitorsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_tenant_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cost_anomaly_monitor_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CostadCostAnomalyMonitorResource()),
						},
					},
				},
			},
		},
	}
}

func readCostadCostAnomalyMonitorsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type CostadCostAnomalyMonitorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_costad.CostAdClient
	Res    *oci_costad.ListCostAnomalyMonitorsResponse
}

func (s *CostadCostAnomalyMonitorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CostadCostAnomalyMonitorsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_costad.ListCostAnomalyMonitorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_costad.CostAnomalyMonitorLifecycleStateEnum(state.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "costad")

	response, err := s.Client.ListCostAnomalyMonitors(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCostAnomalyMonitors(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CostadCostAnomalyMonitorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CostadCostAnomalyMonitorsDataSource-", CostadCostAnomalyMonitorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	costAnomalyMonitor := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CostAnomalyMonitorSummaryToMap(item))
	}
	costAnomalyMonitor["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CostadCostAnomalyMonitorsDataSource().Schema["cost_anomaly_monitor_collection"].Elem.(*schema.Resource).Schema)
		costAnomalyMonitor["items"] = items
	}

	resources = append(resources, costAnomalyMonitor)
	if err := s.D.Set("cost_anomaly_monitor_collection", resources); err != nil {
		return err
	}

	return nil
}
