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

func BudgetCostAnomalyEventResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBudgetCostAnomalyEventWithContext,
		ReadContext:   readBudgetCostAnomalyEventWithContext,
		UpdateContext: updateBudgetCostAnomalyEventWithContext,
		DeleteContext: deleteBudgetCostAnomalyEventWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"cost_anomaly_event_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"feedback_response": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost_anomaly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost_impact": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"cost_monitor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost_monitor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost_monitor_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost_variance_percentage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"root_cause_detail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_resource_filter": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_anomaly_event_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createBudgetCostAnomalyEventWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAnomalyEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readBudgetCostAnomalyEventWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAnomalyEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBudgetCostAnomalyEventWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BudgetCostAnomalyEventResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CostAdClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteBudgetCostAnomalyEventWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type BudgetCostAnomalyEventResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_budget.CostAdClient
	Res                    *oci_budget.CostAnomalyEvent
	DisableNotFoundRetries bool
}

func (s *BudgetCostAnomalyEventResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BudgetCostAnomalyEventResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BudgetCostAnomalyEventResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_budget.CostAnomalyEventLifecycleStateActive),
	}
}

func (s *BudgetCostAnomalyEventResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BudgetCostAnomalyEventResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *BudgetCostAnomalyEventResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_budget.UpdateCostAnomalyEventRequest{}

	if costAnomalyEventId, ok := s.D.GetOkExists("id"); ok {
		tmp := costAnomalyEventId.(string)
		request.CostAnomalyEventId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if feedbackResponse, ok := s.D.GetOkExists("feedback_response"); ok {
		request.FeedbackResponse = oci_budget.CostAnomalyEventFeedbackResponseEnum(feedbackResponse.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.UpdateCostAnomalyEvent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAnomalyEvent
	return nil
}

func (s *BudgetCostAnomalyEventResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_budget.GetCostAnomalyEventRequest{}

	tmp := s.D.Id()
	request.CostAnomalyEventId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.GetCostAnomalyEvent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAnomalyEvent
	return nil
}

func (s *BudgetCostAnomalyEventResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_budget.UpdateCostAnomalyEventRequest{}

	tmp := s.D.Id()
	request.CostAnomalyEventId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if feedbackResponse, ok := s.D.GetOkExists("feedback_response"); ok {
		request.FeedbackResponse = oci_budget.CostAnomalyEventFeedbackResponseEnum(feedbackResponse.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.UpdateCostAnomalyEvent(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAnomalyEvent
	return nil
}

func (s *BudgetCostAnomalyEventResourceCrud) SetData() error {
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

	if s.Res.RootCauseDetail != nil {
		s.D.Set("root_cause_detail", []interface{}{RootCauseDetailToMap(s.Res.RootCauseDetail)})
	} else {
		s.D.Set("root_cause_detail", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetResourceFilter != nil {
		s.D.Set("target_resource_filter", *s.Res.TargetResourceFilter)
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

func CostAnomalyEventSummaryToMap(obj oci_budget.CostAnomalyEventSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CostAnomalyName != nil {
		result["cost_anomaly_name"] = string(*obj.CostAnomalyName)
	}

	if obj.CostImpact != nil {
		result["cost_impact"] = float64(*obj.CostImpact)
	}

	if obj.CostMonitorId != nil {
		result["cost_monitor_id"] = string(*obj.CostMonitorId)
	}

	if obj.CostMonitorName != nil {
		result["cost_monitor_name"] = string(*obj.CostMonitorName)
	}

	//result["cost_monitor_type"] = string(*obj.CostMonitorType)

	if obj.CostVariancePercentage != nil {
		result["cost_variance_percentage"] = float64(*obj.CostVariancePercentage)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	//if obj.RootCauseDetail != nil {
	//	result["root_cause_detail"] = []interface{}{RootCauseDetailToMap(obj.RootCauseDetail)}
	//}
	if obj.RootCauseDetail != nil {
		rootCauseJSON, err := json.Marshal(obj.RootCauseDetail)
		if err != nil {
			log.Printf("[WARN] Unable to marshal root_cause_detail: %v", err)
			result["root_cause_detail"] = ""
		} else {
			result["root_cause_detail"] = string(rootCauseJSON)
		}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetResourceFilter != nil {
		tmp, _ := json.Marshal(obj.TargetResourceFilter)
		result["target_resource_filter"] = string(tmp)
	}

	if obj.TimeAnomalyEventDate != nil {
		result["time_anomaly_event_date"] = obj.TimeAnomalyEventDate.String()
	}

	return result
}

func RootCauseDetailToMap(obj *oci_budget.RootCauseDetail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Value != nil {
		result["value"] = []interface{}{obj.Value}
	}

	return result
}
