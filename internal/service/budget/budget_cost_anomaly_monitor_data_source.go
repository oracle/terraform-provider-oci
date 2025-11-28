// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BudgetCostAnomalyMonitorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cost_anomaly_monitor_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BudgetCostAnomalyMonitorResource(), fieldMap, readSingularBudgetCostAnomalyMonitorWithContext)
}

func readSingularBudgetCostAnomalyMonitorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAnomalyMonitorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BudgetCostAnomalyMonitorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_budget.CostAdClient
	Res    *oci_budget.GetCostAnomalyMonitorResponse
}

func (s *BudgetCostAnomalyMonitorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BudgetCostAnomalyMonitorDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_budget.GetCostAnomalyMonitorRequest{}

	if costAnomalyMonitorId, ok := s.D.GetOkExists("cost_anomaly_monitor_id"); ok {
		tmp := costAnomalyMonitorId.(string)
		request.CostAnomalyMonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "budget")

	response, err := s.Client.GetCostAnomalyMonitor(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BudgetCostAnomalyMonitorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CostAlertSubscriptionMap != nil {
		s.D.Set("cost_alert_subscription_map", []interface{}{CostAlertSubscriptionMapToMap(s.Res.CostAlertSubscriptionMap)})
	} else {
		s.D.Set("cost_alert_subscription_map", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	//if s.Res.TargetResourceFilter != nil {
	//	s.D.Set("target_resource_filter", *s.Res.TargetResourceFilter)
	//}
	if s.Res.TargetResourceFilter != nil {
		targetResourceFilterJSON, err := json.Marshal(s.Res.TargetResourceFilter)
		if err != nil {
			return err
		}
		if err := s.D.Set("target_resource_filter", string(targetResourceFilterJSON)); err != nil {
			return err
		}
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
