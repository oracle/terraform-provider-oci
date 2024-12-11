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

func FleetAppsManagementSchedulerJobJobActivityResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementSchedulerJobJobActivityResources,
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
			"resource_collection": {
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
									"resource_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_id": {
										Type:     schema.TypeString,
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
									"targets": {
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
												"status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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

func readFleetAppsManagementSchedulerJobJobActivityResources(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerJobJobActivityResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementSchedulerJobJobActivityResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListResourcesResponse
}

func (s *FleetAppsManagementSchedulerJobJobActivityResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementSchedulerJobJobActivityResourcesDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListResourcesRequest{}

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

	response, err := s.Client.ListResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementSchedulerJobJobActivityResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementSchedulerJobJobActivityResourcesDataSource-", FleetAppsManagementSchedulerJobJobActivityResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	schedulerJobJobActivityResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceSummaryToMap(item))
	}
	schedulerJobJobActivityResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementSchedulerJobJobActivityResourcesDataSource().Schema["resource_collection"].Elem.(*schema.Resource).Schema)
		schedulerJobJobActivityResource["items"] = items
	}

	resources = append(resources, schedulerJobJobActivityResource)
	if err := s.D.Set("resource_collection", resources); err != nil {
		return err
	}

	return nil
}

func ActivityResourceTargetToMap(obj oci_fleet_apps_management.ActivityResourceTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["status"] = string(obj.Status)

	if obj.TargetName != nil {
		result["target_name"] = string(*obj.TargetName)
	}

	return result
}

func ResourceSummaryToMap(obj oci_fleet_apps_management.ResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.ResourceDisplayName != nil {
		result["resource_display_name"] = string(*obj.ResourceDisplayName)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.Sequence != nil {
		result["sequence"] = string(*obj.Sequence)
	}

	result["status"] = string(obj.Status)

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, ActivityResourceTargetToMap(item))
	}
	result["targets"] = targets

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
