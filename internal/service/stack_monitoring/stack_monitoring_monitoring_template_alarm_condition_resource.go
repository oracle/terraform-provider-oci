// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoringTemplateAlarmConditionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoringTemplateAlarmCondition,
		Read:     readStackMonitoringMonitoringTemplateAlarmCondition,
		Update:   updateStackMonitoringMonitoringTemplateAlarmCondition,
		Delete:   deleteStackMonitoringMonitoringTemplateAlarmCondition,
		Schema: map[string]*schema.Schema{
			// Required
			"condition_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"conditions": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"query": {
							Type:     schema.TypeString,
							Required: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"body": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"should_append_note": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"should_append_url": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"trigger_delay": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"metric_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"monitoring_template_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"composite_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
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
		},
	}
}

func createStackMonitoringMonitoringTemplateAlarmCondition(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateAlarmConditionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoringTemplateAlarmCondition(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateAlarmConditionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringMonitoringTemplateAlarmCondition(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateAlarmConditionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMonitoringTemplateAlarmCondition(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateAlarmConditionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMonitoringTemplateAlarmConditionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.AlarmCondition
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) ID() string {
	return GetMonitoringTemplateAlarmConditionCompositeId(*s.Res.Id, s.D.Get("monitoring_template_id").(string))
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.AlarmConditionLifeCycleStatesCreating),
	}
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.AlarmConditionLifeCycleStatesActive),
	}
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.AlarmConditionLifeCycleStatesDeleted),
	}
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateAlarmConditionRequest{}

	if compositeType, ok := s.D.GetOkExists("composite_type"); ok {
		tmp := compositeType.(string)
		request.CompositeType = &tmp
	}

	if conditionType, ok := s.D.GetOkExists("condition_type"); ok {
		request.ConditionType = oci_stack_monitoring.ConditionTypeEnum(conditionType.(string))
	}

	if conditions, ok := s.D.GetOkExists("conditions"); ok {
		interfaces := conditions.([]interface{})
		tmp := make([]oci_stack_monitoring.Condition, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditions", stateDataIndex)
			converted, err := s.mapToCondition(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("conditions") {
			request.Conditions = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if metricName, ok := s.D.GetOkExists("metric_name"); ok {
		tmp := metricName.(string)
		request.MetricName = &tmp
	}

	if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
		tmp := monitoringTemplateId.(string)
		request.MonitoringTemplateId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateAlarmCondition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AlarmCondition
	return nil
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) Get() error {
	request := oci_stack_monitoring.GetAlarmConditionRequest{}

	alarmConditionId, monitoringTemplateId, err := parseMonitoringTemplateAlarmConditionCompositeId(s.D.Id())
	if err == nil {
		request.AlarmConditionId = &alarmConditionId
		request.MonitoringTemplateId = &monitoringTemplateId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetAlarmCondition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AlarmCondition
	return nil
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) Update() error {
	request := oci_stack_monitoring.UpdateAlarmConditionRequest{}

	alarmConditionId, monitoringTemplateId, err := parseMonitoringTemplateAlarmConditionCompositeId(s.D.Id())
	if err == nil {
		request.AlarmConditionId = &alarmConditionId
		request.MonitoringTemplateId = &monitoringTemplateId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	if compositeType, ok := s.D.GetOkExists("composite_type"); ok {
		tmp := compositeType.(string)
		request.CompositeType = &tmp
	}

	if conditionType, ok := s.D.GetOkExists("condition_type"); ok {
		request.ConditionType = oci_stack_monitoring.ConditionTypeEnum(conditionType.(string))
	}

	if conditions, ok := s.D.GetOkExists("conditions"); ok {
		interfaces := conditions.([]interface{})
		tmp := make([]oci_stack_monitoring.Condition, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditions", stateDataIndex)
			converted, err := s.mapToCondition(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("conditions") {
			request.Conditions = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if metricName, ok := s.D.GetOkExists("metric_name"); ok {
		tmp := metricName.(string)
		request.MetricName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateAlarmCondition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AlarmCondition
	return nil
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteAlarmConditionRequest{}

	if tmp := s.D.Id(); tmp != "" {
		alarmConditionId, monitoringTemplateId, err := parseMonitoringTemplateAlarmConditionCompositeId(s.D.Id())
		if err == nil {
			request.AlarmConditionId = &alarmConditionId
			request.MonitoringTemplateId = &monitoringTemplateId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteAlarmCondition(context.Background(), request)
	return err
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) SetData() error {
	_, monitoringTemplateId, err := parseMonitoringTemplateAlarmConditionCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("monitoring_template_id", &monitoringTemplateId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompositeType != nil {
		s.D.Set("composite_type", *s.Res.CompositeType)
	}

	s.D.Set("condition_type", s.Res.ConditionType)

	conditions := []interface{}{}
	for _, item := range s.Res.Conditions {
		conditions = append(conditions, ConditionToMap(item))
	}
	s.D.Set("conditions", conditions)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MetricName != nil {
		s.D.Set("metric_name", *s.Res.MetricName)
	}

	if s.Res.MonitoringTemplateId != nil {
		s.D.Set("monitoring_template_id", *s.Res.MonitoringTemplateId)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetMonitoringTemplateAlarmConditionCompositeId(alarmConditionId string, monitoringTemplateId string) string {
	alarmConditionId = url.PathEscape(alarmConditionId)
	monitoringTemplateId = url.PathEscape(monitoringTemplateId)
	compositeId := "monitoringTemplates/" + monitoringTemplateId + "/alarmConditions/" + alarmConditionId
	return compositeId
}

func parseMonitoringTemplateAlarmConditionCompositeId(compositeId string) (alarmConditionId string, monitoringTemplateId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("monitoringTemplates/.*/alarmConditions/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	monitoringTemplateId, _ = url.PathUnescape(parts[1])
	alarmConditionId, _ = url.PathUnescape(parts[3])

	return
}

func AlarmConditionSummaryToMap(obj oci_stack_monitoring.AlarmConditionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompositeType != nil {
		result["composite_type"] = string(*obj.CompositeType)
	}

	result["condition_type"] = string(obj.ConditionType)

	conditions := []interface{}{}
	for _, item := range obj.Conditions {
		conditions = append(conditions, ConditionToMap(item))
	}
	result["conditions"] = conditions

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	if obj.MonitoringTemplateId != nil {
		result["monitoring_template_id"] = string(*obj.MonitoringTemplateId)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionResourceCrud) mapToCondition(fieldKeyFormat string) (oci_stack_monitoring.Condition, error) {
	result := oci_stack_monitoring.Condition{}

	if body, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "body")); ok {
		tmp := body.(string)
		result.Body = &tmp
	}

	if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
		tmp := query.(string)
		result.Query = &tmp
	}

	if severity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "severity")); ok {
		result.Severity = oci_stack_monitoring.AlarmConditionSeverityEnum(severity.(string))
	}

	if shouldAppendNote, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_append_note")); ok {
		tmp := shouldAppendNote.(bool)
		result.ShouldAppendNote = &tmp
	}

	if shouldAppendUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_append_url")); ok {
		tmp := shouldAppendUrl.(bool)
		result.ShouldAppendUrl = &tmp
	}

	if triggerDelay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_delay")); ok {
		tmp := triggerDelay.(string)
		result.TriggerDelay = &tmp
	}

	return result, nil
}

func ConditionToMap(obj oci_stack_monitoring.Condition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Body != nil {
		result["body"] = string(*obj.Body)
	}

	if obj.Query != nil {
		result["query"] = string(*obj.Query)
	}

	result["severity"] = string(obj.Severity)

	if obj.ShouldAppendNote != nil {
		result["should_append_note"] = bool(*obj.ShouldAppendNote)
	}

	if obj.ShouldAppendUrl != nil {
		result["should_append_url"] = bool(*obj.ShouldAppendUrl)
	}

	if obj.TriggerDelay != nil {
		result["trigger_delay"] = string(*obj.TriggerDelay)
	}

	return result
}
