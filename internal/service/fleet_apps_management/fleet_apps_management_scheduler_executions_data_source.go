// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementSchedulerExecutionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementSchedulerExecutions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"lifecycle_operation": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runbook_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runbook_version_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_defintion_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_job_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"substate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_scheduled_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_scheduled_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_execution_collection": {
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
									"activity_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"latest_runbook_version_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
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
									"runbook_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"runbook_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"runbook_version_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"scheduler_definition": {
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
												"is_recurring": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"scheduler_job_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_scheduled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readFleetAppsManagementSchedulerExecutions(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerExecutionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementSchedulerExecutionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.ListSchedulerExecutionsResponse
}

func (s *FleetAppsManagementSchedulerExecutionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementSchedulerExecutionsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListSchedulerExecutionsRequest{}

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

	if lifecycleOperation, ok := s.D.GetOkExists("lifecycle_operation"); ok {
		tmp := lifecycleOperation.(string)
		request.LifecycleOperation = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if runbookId, ok := s.D.GetOkExists("runbook_id"); ok {
		tmp := runbookId.(string)
		request.RunbookId = &tmp
	}

	if runbookVersionName, ok := s.D.GetOkExists("runbook_version_name"); ok {
		tmp := runbookVersionName.(string)
		request.RunbookVersionName = &tmp
	}

	if schedulerDefintionId, ok := s.D.GetOkExists("scheduler_defintion_id"); ok {
		tmp := schedulerDefintionId.(string)
		request.SchedulerDefintionId = &tmp
	}

	if schedulerJobId, ok := s.D.GetOkExists("scheduler_job_id"); ok {
		tmp := schedulerJobId.(string)
		request.SchedulerJobId = &tmp
	}

	if substate, ok := s.D.GetOkExists("substate"); ok {
		tmp := substate.(string)
		request.Substate = &tmp
	}

	if timeScheduledGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_scheduled_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduledGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeScheduledGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeScheduledLessThan, ok := s.D.GetOkExists("time_scheduled_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduledLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeScheduledLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListSchedulerExecutions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSchedulerExecutions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementSchedulerExecutionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementSchedulerExecutionsDataSource-", FleetAppsManagementSchedulerExecutionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	schedulerExecution := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SchedulerExecutionSummaryToMap(item))
	}
	schedulerExecution["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementSchedulerExecutionsDataSource().Schema["scheduler_execution_collection"].Elem.(*schema.Resource).Schema)
		schedulerExecution["items"] = items
	}

	resources = append(resources, schedulerExecution)
	if err := s.D.Set("scheduler_execution_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssociatedSchedulerDefinitionToMap(obj *oci_fleet_apps_management.AssociatedSchedulerDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsRecurring != nil {
		result["is_recurring"] = bool(*obj.IsRecurring)
	}

	return result
}

func SchedulerExecutionSummaryToMap(obj oci_fleet_apps_management.SchedulerExecutionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActivityId != nil {
		result["activity_id"] = string(*obj.ActivityId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompartmentName != nil {
		result["compartment_name"] = string(*obj.CompartmentName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LatestRunbookVersionName != nil {
		result["latest_runbook_version_name"] = string(*obj.LatestRunbookVersionName)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ResourceDisplayName != nil {
		result["resource_display_name"] = string(*obj.ResourceDisplayName)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.RunbookDisplayName != nil {
		result["runbook_display_name"] = string(*obj.RunbookDisplayName)
	}

	if obj.RunbookId != nil {
		result["runbook_id"] = string(*obj.RunbookId)
	}

	if obj.RunbookVersionName != nil {
		result["runbook_version_name"] = string(*obj.RunbookVersionName)
	}

	if obj.SchedulerDefinition != nil {
		result["scheduler_definition"] = []interface{}{AssociatedSchedulerDefinitionToMap(obj.SchedulerDefinition)}
	}

	if obj.SchedulerJobId != nil {
		result["scheduler_job_id"] = string(*obj.SchedulerJobId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeScheduled != nil {
		result["time_scheduled"] = obj.TimeScheduled.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
