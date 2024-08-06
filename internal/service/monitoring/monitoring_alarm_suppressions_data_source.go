// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_monitoring "github.com/oracle/oci-go-sdk/v65/monitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MonitoringAlarmSuppressionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMonitoringAlarmSuppressions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"alarm_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_all_suppressions": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alarm_suppression_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MonitoringAlarmSuppressionResource()),
						},
					},
				},
			},
		},
	}
}

func readMonitoringAlarmSuppressions(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmSuppressionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoringClient()

	return tfresource.ReadResource(sync)
}

type MonitoringAlarmSuppressionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.ListAlarmSuppressionsResponse
}

func (s *MonitoringAlarmSuppressionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringAlarmSuppressionsDataSourceCrud) Get() error {
	request := oci_monitoring.ListAlarmSuppressionsRequest{}

	if alarmId, ok := s.D.GetOkExists("alarm_id"); ok {
		tmp := alarmId.(string)
		request.AlarmId = &tmp
	}

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

	if isAllSuppressions, ok := s.D.GetOkExists("is_all_suppressions"); ok {
		tmp := isAllSuppressions.(bool)
		request.IsAllSuppressions = &tmp
	}

	if level, ok := s.D.GetOkExists("level"); ok {
		request.Level = oci_monitoring.AlarmSuppressionLevelEnum(level.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_monitoring.AlarmSuppressionLifecycleStateEnum(state.(string))
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_monitoring.ListAlarmSuppressionsTargetTypeEnum(targetType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "monitoring")

	response, err := s.Client.ListAlarmSuppressions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAlarmSuppressions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MonitoringAlarmSuppressionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MonitoringAlarmSuppressionsDataSource-", MonitoringAlarmSuppressionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	alarmSuppression := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AlarmSuppressionSummaryToMap(item))
	}
	alarmSuppression["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MonitoringAlarmSuppressionsDataSource().Schema["alarm_suppression_collection"].Elem.(*schema.Resource).Schema)
		alarmSuppression["items"] = items
	}

	resources = append(resources, alarmSuppression)
	if err := s.D.Set("alarm_suppression_collection", resources); err != nil {
		return err
	}

	return nil
}
