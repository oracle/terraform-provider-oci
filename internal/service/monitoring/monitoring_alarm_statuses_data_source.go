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

func MonitoringAlarmStatusesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMonitoringAlarmStatuses,
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
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alarm_statuses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"suppression": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_suppress_from": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_suppress_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"timestamp_triggered": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMonitoringAlarmStatuses(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmStatusesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MonitoringClient()

	return tfresource.ReadResource(sync)
}

type MonitoringAlarmStatusesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.ListAlarmsStatusResponse
}

func (s *MonitoringAlarmStatusesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringAlarmStatusesDataSourceCrud) Get() error {
	request := oci_monitoring.ListAlarmsStatusRequest{}

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

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_monitoring.ListAlarmsStatusStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "monitoring")

	response, err := s.Client.ListAlarmsStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAlarmsStatus(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MonitoringAlarmStatusesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MonitoringAlarmStatusesDataSource-", MonitoringAlarmStatusesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		alarmStatus := map[string]interface{}{}

		if r.DisplayName != nil {
			alarmStatus["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			alarmStatus["id"] = *r.Id
		}

		if r.RuleName != nil {
			alarmStatus["rule_name"] = *r.RuleName
		}

		alarmStatus["severity"] = r.Severity

		alarmStatus["status"] = r.Status

		if r.Suppression != nil {
			alarmStatus["suppression"] = []interface{}{SuppressionToMap(r.Suppression)}
		} else {
			alarmStatus["suppression"] = nil
		}

		if r.TimestampTriggered != nil {
			alarmStatus["timestamp_triggered"] = r.TimestampTriggered.String()
		}

		resources = append(resources, alarmStatus)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MonitoringAlarmStatusesDataSource().Schema["alarm_statuses"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("alarm_statuses", resources); err != nil {
		return err
	}

	return nil
}
