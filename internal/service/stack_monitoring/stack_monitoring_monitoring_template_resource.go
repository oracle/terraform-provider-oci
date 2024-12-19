// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoringTemplateResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoringTemplate,
		Read:     readStackMonitoringMonitoringTemplate,
		Update:   updateStackMonitoringMonitoringTemplate,
		Delete:   deleteStackMonitoringMonitoringTemplate,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destinations": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"composite_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
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
			"is_alarms_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_split_notification_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"message_format": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"repeat_notification_duration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"tenant_id": {
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
			"total_alarm_conditions": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"total_applied_alarm_conditions": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func createStackMonitoringMonitoringTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoringTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringMonitoringTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMonitoringTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringMonitoringTemplateResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoringTemplate
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.MonitoringTemplateLifeCycleStatesCreating),
	}
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.MonitoringTemplateLifeCycleStatesActive),
	}
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.MonitoringTemplateLifeCycleStatesDeleted),
	}
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateMonitoringTemplateRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if destinations, ok := s.D.GetOkExists("destinations"); ok {
		interfaces := destinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("destinations") {
			request.Destinations = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAlarmsEnabled, ok := s.D.GetOkExists("is_alarms_enabled"); ok {
		tmp := isAlarmsEnabled.(bool)
		request.IsAlarmsEnabled = &tmp
	}

	if isSplitNotificationEnabled, ok := s.D.GetOkExists("is_split_notification_enabled"); ok {
		tmp := isSplitNotificationEnabled.(bool)
		request.IsSplitNotificationEnabled = &tmp
	}

	if members, ok := s.D.GetOkExists("members"); ok {
		interfaces := members.([]interface{})
		tmp := make([]oci_stack_monitoring.MemberReference, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToMemberReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	if messageFormat, ok := s.D.GetOkExists("message_format"); ok {
		request.MessageFormat = oci_stack_monitoring.MessageFormatEnum(messageFormat.(string))
	}

	if repeatNotificationDuration, ok := s.D.GetOkExists("repeat_notification_duration"); ok {
		tmp := repeatNotificationDuration.(string)
		request.RepeatNotificationDuration = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateMonitoringTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoringTemplate
	return nil
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoringTemplateRequest{}

	tmp := s.D.Id()
	request.MonitoringTemplateId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetMonitoringTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoringTemplate
	return nil
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) Update() error {
	request := oci_stack_monitoring.UpdateMonitoringTemplateRequest{}

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

	if destinations, ok := s.D.GetOkExists("destinations"); ok {
		interfaces := destinations.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("destinations") {
			request.Destinations = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAlarmsEnabled, ok := s.D.GetOkExists("is_alarms_enabled"); ok {
		tmp := isAlarmsEnabled.(bool)
		request.IsAlarmsEnabled = &tmp
	}

	if isSplitNotificationEnabled, ok := s.D.GetOkExists("is_split_notification_enabled"); ok {
		tmp := isSplitNotificationEnabled.(bool)
		request.IsSplitNotificationEnabled = &tmp
	}

	if members, ok := s.D.GetOkExists("members"); ok {
		interfaces := members.([]interface{})
		tmp := make([]oci_stack_monitoring.MemberReference, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "members", stateDataIndex)
			converted, err := s.mapToMemberReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("members") {
			request.Members = tmp
		}
	}

	if messageFormat, ok := s.D.GetOkExists("message_format"); ok {
		request.MessageFormat = oci_stack_monitoring.MessageFormatEnum(messageFormat.(string))
	}

	tmp := s.D.Id()
	request.MonitoringTemplateId = &tmp

	if repeatNotificationDuration, ok := s.D.GetOkExists("repeat_notification_duration"); ok {
		tmp := repeatNotificationDuration.(string)
		request.RepeatNotificationDuration = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateMonitoringTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoringTemplate
	return nil
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteMonitoringTemplateRequest{}

	tmp := s.D.Id()
	request.MonitoringTemplateId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteMonitoringTemplate(context.Background(), request)
	return err
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("destinations", s.Res.Destinations)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAlarmsEnabled != nil {
		s.D.Set("is_alarms_enabled", *s.Res.IsAlarmsEnabled)
	}

	if s.Res.IsSplitNotificationEnabled != nil {
		s.D.Set("is_split_notification_enabled", *s.Res.IsSplitNotificationEnabled)
	}

	members := []interface{}{}
	for _, item := range s.Res.Members {
		members = append(members, MemberReferenceToMap(item))
	}
	s.D.Set("members", members)

	s.D.Set("message_format", s.Res.MessageFormat)

	if s.Res.RepeatNotificationDuration != nil {
		s.D.Set("repeat_notification_duration", *s.Res.RepeatNotificationDuration)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalAlarmConditions != nil {
		s.D.Set("total_alarm_conditions", *s.Res.TotalAlarmConditions)
	}

	if s.Res.TotalAppliedAlarmConditions != nil {
		s.D.Set("total_applied_alarm_conditions", *s.Res.TotalAppliedAlarmConditions)
	}

	return nil
}

func (s *StackMonitoringMonitoringTemplateResourceCrud) mapToMemberReference(fieldKeyFormat string) (oci_stack_monitoring.MemberReference, error) {
	result := oci_stack_monitoring.MemberReference{}

	if compositeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "composite_type")); ok {
		tmp := compositeType.(string)
		if len(tmp) != 0 {
			result.CompositeType = &tmp
		}
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_stack_monitoring.MemberReferenceTypeEnum(type_.(string))
	}

	return result, nil
}

func MemberReferenceToMap(obj oci_stack_monitoring.MemberReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompositeType != nil && len(string(*obj.CompositeType)) != 0 {
		result["composite_type"] = string(*obj.CompositeType)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["type"] = string(obj.Type)

	return result
}

func MonitoringTemplateSummaryToMap(obj oci_stack_monitoring.MonitoringTemplateSummary) map[string]interface{} {
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

	result["destinations"] = obj.Destinations

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	members := []interface{}{}
	for _, item := range obj.Members {
		members = append(members, MemberReferenceToMap(item))
	}
	result["members"] = members

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TotalAlarmConditions != nil {
		result["total_alarm_conditions"] = float32(*obj.TotalAlarmConditions)
	}

	if obj.TotalAppliedAlarmConditions != nil {
		result["total_applied_alarm_conditions"] = float32(*obj.TotalAppliedAlarmConditions)
	}

	return result
}
