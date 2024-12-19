// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoringTemplateDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["monitoring_template_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringMonitoringTemplateResource(), fieldMap, readSingularStackMonitoringMonitoringTemplate)
}

func readSingularStackMonitoringMonitoringTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoringTemplateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetMonitoringTemplateResponse
}

func (s *StackMonitoringMonitoringTemplateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoringTemplateDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoringTemplateRequest{}

	if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
		tmp := monitoringTemplateId.(string)
		request.MonitoringTemplateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetMonitoringTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringMonitoringTemplateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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
