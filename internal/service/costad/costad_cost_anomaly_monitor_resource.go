// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package costad

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_costad "github.com/oracle/oci-go-sdk/v65/costad"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CostadCostAnomalyMonitorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createCostadCostAnomalyMonitorWithContext,
		ReadContext:   readCostadCostAnomalyMonitorWithContext,
		UpdateContext: updateCostadCostAnomalyMonitorWithContext,
		DeleteContext: deleteCostadCostAnomalyMonitorWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_resource_filter": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},

			// Optional
			"cost_alert_subscription_map": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"cost_alert_subscription_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"threshold_absolute_value": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"threshold_relative_percent": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
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
			"lifecycle_details": {
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
			"time_created": {
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
		},
	}
}

func createCostadCostAnomalyMonitorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readCostadCostAnomalyMonitorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateCostadCostAnomalyMonitorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteCostadCostAnomalyMonitorWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAnomalyMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type CostadCostAnomalyMonitorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_costad.CostAdClient
	Res                    *oci_costad.CostAnomalyMonitor
	DisableNotFoundRetries bool
}

func (s *CostadCostAnomalyMonitorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CostadCostAnomalyMonitorResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *CostadCostAnomalyMonitorResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_costad.CostAnomalyMonitorLifecycleStateActive),
	}
}

func (s *CostadCostAnomalyMonitorResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CostadCostAnomalyMonitorResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_costad.CostAnomalyMonitorLifecycleStateDeleted),
	}
}

func (s *CostadCostAnomalyMonitorResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_costad.CreateCostAnomalyMonitorRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if costAlertSubscriptionMap, ok := s.D.GetOkExists("cost_alert_subscription_map"); ok {
		if tmpList := costAlertSubscriptionMap.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cost_alert_subscription_map", 0)
			tmp, err := s.mapToCostAlertSubscriptionMap(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CostAlertSubscriptionMap = &tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if targetResourceFilter, ok := s.D.GetOkExists("target_resource_filter"); ok {
		tmp := targetResourceFilter.(string)
		var targetResourceFilterObj oci_costad.TargetResourceFilter
		err := json.Unmarshal([]byte(tmp), &targetResourceFilterObj)
		if err != nil {
			return err
		}
		request.TargetResourceFilter = &targetResourceFilterObj
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.CreateCostAnomalyMonitor(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAnomalyMonitor
	return nil
}

func (s *CostadCostAnomalyMonitorResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_costad.GetCostAnomalyMonitorRequest{}

	tmp := s.D.Id()
	request.CostAnomalyMonitorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.GetCostAnomalyMonitor(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAnomalyMonitor
	return nil
}

func (s *CostadCostAnomalyMonitorResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_costad.UpdateCostAnomalyMonitorRequest{}

	if costAlertSubscriptionMap, ok := s.D.GetOkExists("cost_alert_subscription_map"); ok {
		if tmpList := costAlertSubscriptionMap.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "cost_alert_subscription_map", 0)
			tmp, err := s.mapToCostAlertSubscriptionMap(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CostAlertSubscriptionMap = &tmp
		}
	}

	tmp := s.D.Id()
	request.CostAnomalyMonitorId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	response, err := s.Client.UpdateCostAnomalyMonitor(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.CostAnomalyMonitor
	return nil
}

func (s *CostadCostAnomalyMonitorResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_costad.DeleteCostAnomalyMonitorRequest{}

	tmp := s.D.Id()
	request.CostAnomalyMonitorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "costad")

	_, err := s.Client.DeleteCostAnomalyMonitor(ctx, request)
	return err
}

func (s *CostadCostAnomalyMonitorResourceCrud) SetData() error {
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

	if s.Res.TargetResourceFilter != nil {
		targetResourceFilterJSON, err := json.Marshal(s.Res.TargetResourceFilter)
		if err != nil {
			return err
		}
		s.D.Set("target_resource_filter", string(targetResourceFilterJSON))
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

func (s *CostadCostAnomalyMonitorResourceCrud) mapToCostAlertSubscriptionMap(fieldKeyFormat string) (oci_costad.CostAlertSubscriptionMap, error) {
	result := oci_costad.CostAlertSubscriptionMap{}

	if costAlertSubscriptionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cost_alert_subscription_id")); ok {
		tmp := costAlertSubscriptionId.(string)
		result.CostAlertSubscriptionId = &tmp
	}

	if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
		result.Operator = oci_costad.CostAlertSubscriptionMapOperatorEnum(operator.(string))
	}

	if thresholdAbsoluteValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold_absolute_value")); ok {
		tmp := thresholdAbsoluteValue.(int)
		result.ThresholdAbsoluteValue = &tmp
	}

	if thresholdRelativePercent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold_relative_percent")); ok {
		tmp := thresholdRelativePercent.(int)
		result.ThresholdRelativePercent = &tmp
	}

	return result, nil
}

func CostAlertSubscriptionMapToMap(obj *oci_costad.CostAlertSubscriptionMap) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CostAlertSubscriptionId != nil {
		result["cost_alert_subscription_id"] = string(*obj.CostAlertSubscriptionId)
	}

	result["operator"] = string(obj.Operator)

	if obj.ThresholdAbsoluteValue != nil {
		result["threshold_absolute_value"] = int(*obj.ThresholdAbsoluteValue)
	}

	if obj.ThresholdRelativePercent != nil {
		result["threshold_relative_percent"] = int(*obj.ThresholdRelativePercent)
	}

	return result
}

func CostAnomalyMonitorSummaryToMap(obj oci_costad.CostAnomalyMonitorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetResourceFilter != nil {
		tmp, _ := json.Marshal(obj.TargetResourceFilter)
		result["target_resource_filter"] = string(tmp)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}
