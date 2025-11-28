// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BudgetCostAnomalyEventDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cost_anomaly_event_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BudgetCostAnomalyEventResource(), fieldMap, readSingularBudgetCostAnomalyEventWithContext)
}

func readSingularBudgetCostAnomalyEventWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAnomalyEventDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BudgetCostAnomalyEventDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_budget.CostAdClient
	Res    *oci_budget.GetCostAnomalyEventResponse
}

func (s *BudgetCostAnomalyEventDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BudgetCostAnomalyEventDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_budget.GetCostAnomalyEventRequest{}

	if costAnomalyEventId, ok := s.D.GetOkExists("cost_anomaly_event_id"); ok {
		tmp := costAnomalyEventId.(string)
		request.CostAnomalyEventId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "budget")

	response, err := s.Client.GetCostAnomalyEvent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BudgetCostAnomalyEventDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CostAnomalyName != nil {
		s.D.Set("cost_anomaly_name", *s.Res.CostAnomalyName)
	}

	if s.Res.CostImpact != nil {
		s.D.Set("cost_impact", *s.Res.CostImpact)
	}

	if s.Res.CostMonitorId != nil {
		s.D.Set("cost_monitor_id", *s.Res.CostMonitorId)
	}

	if s.Res.CostMonitorName != nil {
		s.D.Set("cost_monitor_name", *s.Res.CostMonitorName)
	}

	s.D.Set("cost_monitor_type", s.Res.CostMonitorType)

	if s.Res.CostVariancePercentage != nil {
		s.D.Set("cost_variance_percentage", *s.Res.CostVariancePercentage)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("feedback_response", s.Res.FeedbackResponse)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	//if s.Res.RootCauseDetail != nil {
	//	s.D.Set("root_cause_detail", []interface{}{RootCauseDetailToMap(s.Res.RootCauseDetail)})
	//} else {
	//	s.D.Set("root_cause_detail", nil)
	//}
	if s.Res.RootCauseDetail != nil {
		rootCauseJSON, err := json.Marshal(s.Res.RootCauseDetail)
		if err != nil {
			log.Printf("[WARN] Unable to marshal root_cause_detail: %v", err)
			s.D.Set("root_cause_detail", "")
		} else {
			s.D.Set("root_cause_detail", string(rootCauseJSON))
		}
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
			log.Printf("[WARN] Unable to marshal target_resource_filter: %v", err)
			s.D.Set("target_resource_filter", "")
		} else {
			s.D.Set("target_resource_filter", string(targetResourceFilterJSON))
		}
	}

	if s.Res.TimeAnomalyEventDate != nil {
		s.D.Set("time_anomaly_event_date", s.Res.TimeAnomalyEventDate.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
