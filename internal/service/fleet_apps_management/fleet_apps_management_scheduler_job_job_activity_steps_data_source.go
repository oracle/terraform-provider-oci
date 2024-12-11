// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementSchedulerJobJobActivityStepsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementSchedulerJobJobActivitySteps,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"job_activity_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_task_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sequence": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"step_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"step_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
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
									"is_rollback_task": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sequence": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"step_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"task_record_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementSchedulerJobJobActivitySteps(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerJobJobActivityStepsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementSchedulerJobJobActivityStepsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListStepsResponse
}

func (s *FleetAppsManagementSchedulerJobJobActivityStepsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementSchedulerJobJobActivityStepsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListStepsRequest{}

	if jobActivityId, ok := s.D.GetOkExists("job_activity_id"); ok {
		tmp := jobActivityId.(string)
		request.JobActivityId = &tmp
	}

	if resourceTaskId, ok := s.D.GetOkExists("resource_task_id"); ok {
		tmp := resourceTaskId.(string)
		request.ResourceTaskId = &tmp
	}

	if schedulerJobId, ok := s.D.GetOkExists("scheduler_job_id"); ok {
		tmp := schedulerJobId.(string)
		request.SchedulerJobId = &tmp
	}

	if sequence, ok := s.D.GetOkExists("sequence"); ok {
		tmp := sequence.(string)
		request.Sequence = &tmp
	}

	if stepName, ok := s.D.GetOkExists("step_name"); ok {
		tmp := stepName.(string)
		request.StepName = &tmp
	}

	if targetName, ok := s.D.GetOkExists("target_name"); ok {
		tmp := targetName.(string)
		request.TargetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListSteps(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSteps(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementSchedulerJobJobActivityStepsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementSchedulerJobJobActivityStepsDataSource-", FleetAppsManagementSchedulerJobJobActivityStepsDataSource(), s.D))
	resources := []map[string]interface{}{}
	schedulerJobJobActivityStep := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, StepSummaryToMap(item))
	}
	schedulerJobJobActivityStep["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementSchedulerJobJobActivityStepsDataSource().Schema["step_collection"].Elem.(*schema.Resource).Schema)
		schedulerJobJobActivityStep["items"] = items
	}

	resources = append(resources, schedulerJobJobActivityStep)
	if err := s.D.Set("step_collection", resources); err != nil {
		return err
	}

	return nil
}

func StepSummaryToMap(obj oci_fleet_apps_management.StepSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsRollbackTask != nil {
		result["is_rollback_task"] = bool(*obj.IsRollbackTask)
	}

	if obj.Sequence != nil {
		result["sequence"] = string(*obj.Sequence)
	}

	result["status"] = string(obj.Status)

	if obj.StepName != nil {
		result["step_name"] = string(*obj.StepName)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TaskRecordId != nil {
		result["task_record_id"] = string(*obj.TaskRecordId)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
