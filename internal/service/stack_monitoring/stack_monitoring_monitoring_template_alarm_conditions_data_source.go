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

func StackMonitoringMonitoringTemplateAlarmConditionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringMonitoringTemplateAlarmConditions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"alarm_condition_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"monitoring_template_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alarm_condition_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(StackMonitoringMonitoringTemplateAlarmConditionResource()),
						},
					},
				},
			},
		},
	}
}

func readStackMonitoringMonitoringTemplateAlarmConditions(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplateAlarmConditionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoringTemplateAlarmConditionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListAlarmConditionsResponse
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionsDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListAlarmConditionsRequest{}

	if metricName, ok := s.D.GetOkExists("metric_name"); ok {
		interfaces := metricName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("metric_name") {
			request.MetricName = tmp
		}
	}

	if monitoringTemplateId, ok := s.D.GetOkExists("monitoring_template_id"); ok {
		tmp := monitoringTemplateId.(string)
		request.MonitoringTemplateId = &tmp
	}

	if resourceTypes, ok := s.D.GetOkExists("resource_types"); ok {
		interfaces := resourceTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_types") {
			request.ResourceTypes = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_stack_monitoring.ListAlarmConditionsLifecycleStateEnum(state.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_stack_monitoring.ListAlarmConditionsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListAlarmConditions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAlarmConditions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringMonitoringTemplateAlarmConditionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringMonitoringTemplateAlarmConditionsDataSource-", StackMonitoringMonitoringTemplateAlarmConditionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitoringTemplateAlarmCondition := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AlarmConditionSummaryToMap(item))
	}
	monitoringTemplateAlarmCondition["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringMonitoringTemplateAlarmConditionsDataSource().Schema["alarm_condition_collection"].Elem.(*schema.Resource).Schema)
		monitoringTemplateAlarmCondition["items"] = items
	}

	resources = append(resources, monitoringTemplateAlarmCondition)
	if err := s.D.Set("alarm_condition_collection", resources); err != nil {
		return err
	}

	return nil
}
