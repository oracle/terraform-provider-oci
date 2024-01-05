// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package monitoring

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_monitoring "github.com/oracle/oci-go-sdk/v65/monitoring"
)

func MonitoringAlarmDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["alarm_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MonitoringAlarmResource(), fieldMap, readSingularMonitoringAlarm)
}

func readSingularMonitoringAlarm(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoringClient()

	return tfresource.ReadResource(sync)
}

type MonitoringAlarmDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.GetAlarmResponse
}

func (s *MonitoringAlarmDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringAlarmDataSourceCrud) Get() error {
	request := oci_monitoring.GetAlarmRequest{}

	if alarmId, ok := s.D.GetOkExists("alarm_id"); ok {
		tmp := alarmId.(string)
		request.AlarmId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "monitoring")

	response, err := s.Client.GetAlarm(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MonitoringAlarmDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Body != nil {
		s.D.Set("body", *s.Res.Body)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("destinations", s.Res.Destinations)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.IsNotificationsPerMetricDimensionEnabled != nil {
		s.D.Set("is_notifications_per_metric_dimension_enabled", *s.Res.IsNotificationsPerMetricDimensionEnabled)
	}

	s.D.Set("message_format", s.Res.MessageFormat)

	if s.Res.MetricCompartmentId != nil {
		s.D.Set("metric_compartment_id", *s.Res.MetricCompartmentId)
	}

	if s.Res.MetricCompartmentIdInSubtree != nil {
		s.D.Set("metric_compartment_id_in_subtree", *s.Res.MetricCompartmentIdInSubtree)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.PendingDuration != nil {
		s.D.Set("pending_duration", *s.Res.PendingDuration)
	}

	if s.Res.Query != nil {
		s.D.Set("query", *s.Res.Query)
	}

	if s.Res.RepeatNotificationDuration != nil {
		s.D.Set("repeat_notification_duration", *s.Res.RepeatNotificationDuration)
	}

	if s.Res.Resolution != nil {
		s.D.Set("resolution", *s.Res.Resolution)
	}

	if s.Res.ResourceGroup != nil {
		s.D.Set("resource_group", *s.Res.ResourceGroup)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Suppression != nil {
		s.D.Set("suppression", []interface{}{SuppressionToMap(s.Res.Suppression)})
	} else {
		s.D.Set("suppression", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
