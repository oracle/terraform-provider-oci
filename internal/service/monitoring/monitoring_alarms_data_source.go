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

func MonitoringAlarmsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMonitoringAlarms,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alarms": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(MonitoringAlarmResource()),
			},
		},
	}
}

func readMonitoringAlarms(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoringClient()

	return tfresource.ReadResource(sync)
}

type MonitoringAlarmsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.ListAlarmsResponse
}

func (s *MonitoringAlarmsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringAlarmsDataSourceCrud) Get() error {
	request := oci_monitoring.ListAlarmsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_monitoring.AlarmLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "monitoring")

	response, err := s.Client.ListAlarms(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAlarms(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MonitoringAlarmsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MonitoringAlarmsDataSource-", MonitoringAlarmsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		alarm := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			alarm["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		alarm["destinations"] = r.Destinations

		if r.DisplayName != nil {
			alarm["display_name"] = *r.DisplayName
		}

		alarm["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			alarm["id"] = *r.Id
		}

		if r.IsEnabled != nil {
			alarm["is_enabled"] = *r.IsEnabled
		}

		if r.IsNotificationsPerMetricDimensionEnabled != nil {
			alarm["is_notifications_per_metric_dimension_enabled"] = *r.IsNotificationsPerMetricDimensionEnabled
		}

		if r.MetricCompartmentId != nil {
			alarm["metric_compartment_id"] = *r.MetricCompartmentId
		}

		if r.Namespace != nil {
			alarm["namespace"] = *r.Namespace
		}

		if r.NotificationVersion != nil {
			alarm["notification_version"] = *r.NotificationVersion
		}

		overrides := []interface{}{}
		for _, item := range r.Overrides {
			overrides = append(overrides, AlarmOverrideToMap(item))
		}
		alarm["overrides"] = overrides

		if r.Query != nil {
			alarm["query"] = *r.Query
		}

		if r.RuleName != nil {
			alarm["rule_name"] = *r.RuleName
		}

		alarm["severity"] = r.Severity

		alarm["state"] = r.LifecycleState

		if r.Suppression != nil {
			alarm["suppression"] = []interface{}{SuppressionToMap(r.Suppression)}
		} else {
			alarm["suppression"] = nil
		}

		resources = append(resources, alarm)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MonitoringAlarmsDataSource().Schema["alarms"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("alarms", resources); err != nil {
		return err
	}

	return nil
}
