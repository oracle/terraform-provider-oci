// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoringTemplateAlarmConditionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["alarm_condition_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["monitoring_template_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringMonitoringTemplateAlarmConditionResource(), fieldMap, readSingularStackMonitoringMonitoringTemplateAlarmCondition)
}

func readSingularStackMonitoringMonitoringTemplateAlarmCondition(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateAlarmConditionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoringTemplateAlarmConditionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetAlarmConditionResponse
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetAlarmConditionRequest{}

	if compositeId, ok := s.D.GetOkExists("alarm_condition_id"); ok {
		compositeId := compositeId.(string)
		alarmConditionId, monitoringTemplateId, err := parseMonitoringTemplateAlarmConditionCompositeId(compositeId)
		if err == nil {
			request.AlarmConditionId = &alarmConditionId
			request.MonitoringTemplateId = &monitoringTemplateId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetAlarmCondition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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
