// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package monitoring

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_monitoring "github.com/oracle/oci-go-sdk/v65/monitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MonitoringAlarmSuppressionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["alarm_suppression_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MonitoringAlarmSuppressionResource(), fieldMap, readSingularMonitoringAlarmSuppression)
}

func readSingularMonitoringAlarmSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmSuppressionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoringClient()

	return tfresource.ReadResource(sync)
}

type MonitoringAlarmSuppressionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.GetAlarmSuppressionResponse
}

func (s *MonitoringAlarmSuppressionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringAlarmSuppressionDataSourceCrud) Get() error {
	request := oci_monitoring.GetAlarmSuppressionRequest{}

	if alarmSuppressionId, ok := s.D.GetOkExists("alarm_suppression_id"); ok {
		tmp := alarmSuppressionId.(string)
		request.AlarmSuppressionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "monitoring")

	response, err := s.Client.GetAlarmSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MonitoringAlarmSuppressionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AlarmSuppressionTarget != nil {
		alarmSuppressionTargetArray := []interface{}{}
		if alarmSuppressionTargetMap := AlarmSuppressionTargetToMap(&s.Res.AlarmSuppressionTarget); alarmSuppressionTargetMap != nil {
			alarmSuppressionTargetArray = append(alarmSuppressionTargetArray, alarmSuppressionTargetMap)
		}
		s.D.Set("alarm_suppression_target", alarmSuppressionTargetArray)
	} else {
		s.D.Set("alarm_suppression_target", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("dimensions", s.Res.Dimensions)
	s.D.Set("dimensions", s.Res.Dimensions)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("level", s.Res.Level)

	s.D.Set("state", s.Res.LifecycleState)

	suppressionConditions := []interface{}{}
	for _, item := range s.Res.SuppressionConditions {
		suppressionConditions = append(suppressionConditions, SuppressionConditionToMap(item))
	}
	s.D.Set("suppression_conditions", suppressionConditions)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeSuppressFrom != nil {
		s.D.Set("time_suppress_from", s.Res.TimeSuppressFrom.Format(time.RFC3339Nano))
	}

	if s.Res.TimeSuppressUntil != nil {
		s.D.Set("time_suppress_until", s.Res.TimeSuppressUntil.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
