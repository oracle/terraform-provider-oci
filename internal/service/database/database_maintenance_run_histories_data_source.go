// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMaintenanceRunHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMaintenanceRunHistories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"maintenance_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_run_histories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"current_execution_window": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_servers_history_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"db_server_patching_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"estimated_patch_duration": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"patching_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_patching_ended": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_patching_started": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"granular_maintenance_history": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"execution_actions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"action_members": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"estimated_time_in_mins": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"member_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"member_order": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"status": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"total_time_taken_in_mins": {
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"action_params": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"action_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"defined_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"estimated_time_in_mins": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"execution_action_order": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"execution_window_id": {
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
												"lifecycle_details": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"lifecycle_substate": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_created": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_updated": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"total_time_taken_in_mins": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"execution_window": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"defined_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"estimated_time_in_mins": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"execution_resource_id": {
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
												"is_enforced_duration": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"lifecycle_details": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"lifecycle_substate": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"state": {
													Type:     schema.TypeString,
													Computed: true,
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
												"total_time_taken_in_mins": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"window_duration_in_mins": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"window_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_run_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"current_custom_action_timeout_in_mins": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"current_patching_component": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_action_timeout_in_mins": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"database_software_image_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"estimated_component_patching_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"estimated_patching_time": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"estimated_db_server_patching_time": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"estimated_network_switches_patching_time": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"estimated_storage_server_patching_time": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_estimated_patching_time": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_custom_action_timeout_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_dst_file_update_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_maintenance_run_granular": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"maintenance_subtype": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"maintenance_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patch_failure_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"patch_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patching_end_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patching_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patching_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"patching_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"peer_maintenance_run_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"peer_maintenance_run_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_db_server_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_storage_server_version": {
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
									"total_time_taken_in_mins": {
										Type:     schema.TypeInt,
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

func readDatabaseMaintenanceRunHistories(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMaintenanceRunHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListMaintenanceRunHistoryResponse
}

func (s *DatabaseMaintenanceRunHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMaintenanceRunHistoriesDataSourceCrud) Get() error {
	request := oci_database.ListMaintenanceRunHistoryRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if maintenanceType, ok := s.D.GetOkExists("maintenance_type"); ok {
		request.MaintenanceType = oci_database.MaintenanceRunSummaryMaintenanceTypeEnum(maintenanceType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.MaintenanceRunSummaryLifecycleStateEnum(state.(string))
	}

	if targetResourceId, ok := s.D.GetOkExists("target_resource_id"); ok {
		tmp := targetResourceId.(string)
		request.TargetResourceId = &tmp
	}

	if targetResourceType, ok := s.D.GetOkExists("target_resource_type"); ok {
		request.TargetResourceType = oci_database.MaintenanceRunSummaryTargetResourceTypeEnum(targetResourceType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListMaintenanceRunHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaintenanceRunHistory(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMaintenanceRunHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMaintenanceRunHistoriesDataSource-", DatabaseMaintenanceRunHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		maintenanceRunHistory := map[string]interface{}{}

		if r.CurrentExecutionWindow != nil {
			maintenanceRunHistory["current_execution_window"] = *r.CurrentExecutionWindow
		}

		dbServersHistoryDetails := []interface{}{}
		for _, item := range r.DbServersHistoryDetails {
			dbServersHistoryDetails = append(dbServersHistoryDetails, DbServerHistorySummaryToMap(item))
		}
		maintenanceRunHistory["db_servers_history_details"] = dbServersHistoryDetails

		granularMaintenanceHistory := []interface{}{}
		for _, item := range r.GranularMaintenanceHistory {
			granularMaintenanceHistory = append(granularMaintenanceHistory, GranularMaintenanceHistoryDetailsToMap(item))
		}
		maintenanceRunHistory["granular_maintenance_history"] = granularMaintenanceHistory

		if r.Id != nil {
			maintenanceRunHistory["id"] = *r.Id
		}

		if r.MaintenanceRunDetails != nil {
			maintenanceRunHistory["maintenance_run_details"] = []interface{}{MaintenanceRunSummaryToMap(r.MaintenanceRunDetails)}
		} else {
			maintenanceRunHistory["maintenance_run_details"] = nil
		}

		resources = append(resources, maintenanceRunHistory)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseMaintenanceRunHistoriesDataSource().Schema["maintenance_run_histories"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("maintenance_run_histories", resources); err != nil {
		return err
	}

	return nil
}

func DbServerHistorySummaryToMap(obj oci_database.DbServerHistorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbServerPatchingDetails != nil {
		result["db_server_patching_details"] = []interface{}{DbServerPatchingDetailToMap(obj.DbServerPatchingDetails)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func DbServerPatchingDetailToMap(obj *oci_database.DbServerPatchingDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EstimatedPatchDuration != nil {
		result["estimated_patch_duration"] = int(*obj.EstimatedPatchDuration)
	}

	result["patching_status"] = string(obj.PatchingStatus)

	if obj.TimePatchingEnded != nil {
		result["time_patching_ended"] = obj.TimePatchingEnded.String()
	}

	if obj.TimePatchingStarted != nil {
		result["time_patching_started"] = obj.TimePatchingStarted.String()
	}

	return result
}

func EstdPatchingTimeToMap(obj *oci_database.EstimatedPatchingTime) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EstimatedDbServerPatchingTime != nil {
		result["estimated_db_server_patching_time"] = int(*obj.EstimatedDbServerPatchingTime)
	}

	if obj.EstimatedNetworkSwitchesPatchingTime != nil {
		result["estimated_network_switches_patching_time"] = int(*obj.EstimatedNetworkSwitchesPatchingTime)
	}

	if obj.EstimatedStorageServerPatchingTime != nil {
		result["estimated_storage_server_patching_time"] = int(*obj.EstimatedStorageServerPatchingTime)
	}

	if obj.TotalEstimatedPatchingTime != nil {
		result["total_estimated_patching_time"] = int(*obj.TotalEstimatedPatchingTime)
	}

	return result
}

func ExecutionActionToMap(obj oci_database.ExecutionAction) map[string]interface{} {
	result := map[string]interface{}{}

	actionMembers := []interface{}{}
	for _, item := range obj.ActionMembers {
		actionMembers = append(actionMembers, ExecutionActionMemberToMap(item))
	}
	result["action_members"] = actionMembers

	result["action_params"] = obj.ActionParams

	result["action_type"] = string(obj.ActionType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EstimatedTimeInMins != nil {
		result["estimated_time_in_mins"] = int(*obj.EstimatedTimeInMins)
	}

	if obj.ExecutionActionOrder != nil {
		result["execution_action_order"] = int(*obj.ExecutionActionOrder)
	}

	if obj.ExecutionWindowId != nil {
		result["execution_window_id"] = string(*obj.ExecutionWindowId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["lifecycle_substate"] = string(obj.LifecycleSubstate)

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TotalTimeTakenInMins != nil {
		result["total_time_taken_in_mins"] = int(*obj.TotalTimeTakenInMins)
	}

	return result
}

func ExecutionWindowToMap(obj *oci_database.ExecutionWindow) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EstimatedTimeInMins != nil {
		result["estimated_time_in_mins"] = int(*obj.EstimatedTimeInMins)
	}

	if obj.ExecutionResourceId != nil {
		result["execution_resource_id"] = string(*obj.ExecutionResourceId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEnforcedDuration != nil {
		result["is_enforced_duration"] = bool(*obj.IsEnforcedDuration)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["lifecycle_substate"] = string(obj.LifecycleSubstate)

	result["state"] = string(obj.LifecycleState)

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

	if obj.TotalTimeTakenInMins != nil {
		result["total_time_taken_in_mins"] = int(*obj.TotalTimeTakenInMins)
	}

	if obj.WindowDurationInMins != nil {
		result["window_duration_in_mins"] = int(*obj.WindowDurationInMins)
	}

	result["window_type"] = string(obj.WindowType)

	return result
}

func GranularMaintenanceHistoryDetailsToMap(obj oci_database.GranularMaintenanceHistoryDetails) map[string]interface{} {
	result := map[string]interface{}{}

	executionActions := []interface{}{}
	for _, item := range obj.ExecutionActions {
		executionActions = append(executionActions, ExecutionActionToMap(item))
	}
	result["execution_actions"] = executionActions

	if obj.ExecutionWindow != nil {
		result["execution_window"] = []interface{}{ExecutionWindowToMap(obj.ExecutionWindow)}
	}

	return result
}

func MaintenanceRunSummaryToMap(obj *oci_database.MaintenanceRunSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CurrentCustomActionTimeoutInMins != nil {
		result["current_custom_action_timeout_in_mins"] = int(*obj.CurrentCustomActionTimeoutInMins)
	}

	if obj.CurrentPatchingComponent != nil {
		result["current_patching_component"] = string(*obj.CurrentPatchingComponent)
	}

	if obj.CustomActionTimeoutInMins != nil {
		result["custom_action_timeout_in_mins"] = int(*obj.CustomActionTimeoutInMins)
	}

	if obj.DatabaseSoftwareImageId != nil {
		result["database_software_image_id"] = string(*obj.DatabaseSoftwareImageId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EstimatedComponentPatchingStartTime != nil {
		result["estimated_component_patching_start_time"] = obj.EstimatedComponentPatchingStartTime.String()
	}

	if obj.EstimatedPatchingTime != nil {
		result["estimated_patching_time"] = []interface{}{EstdPatchingTimeToMap(obj.EstimatedPatchingTime)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCustomActionTimeoutEnabled != nil {
		result["is_custom_action_timeout_enabled"] = bool(*obj.IsCustomActionTimeoutEnabled)
	}

	if obj.IsDstFileUpdateEnabled != nil {
		result["is_dst_file_update_enabled"] = bool(*obj.IsDstFileUpdateEnabled)
	}

	if obj.IsMaintenanceRunGranular != nil {
		result["is_maintenance_run_granular"] = bool(*obj.IsMaintenanceRunGranular)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["maintenance_subtype"] = string(obj.MaintenanceSubtype)

	result["maintenance_type"] = string(obj.MaintenanceType)

	if obj.PatchFailureCount != nil {
		result["patch_failure_count"] = int(*obj.PatchFailureCount)
	}

	if obj.PatchId != nil {
		result["patch_id"] = string(*obj.PatchId)
	}

	if obj.PatchingEndTime != nil {
		result["patching_end_time"] = obj.PatchingEndTime.String()
	}

	result["patching_mode"] = string(obj.PatchingMode)

	if obj.PatchingStartTime != nil {
		result["patching_start_time"] = obj.PatchingStartTime.String()
	}

	result["patching_status"] = string(obj.PatchingStatus)

	if obj.PeerMaintenanceRunId != nil {
		result["peer_maintenance_run_id"] = string(*obj.PeerMaintenanceRunId)
	}

	result["peer_maintenance_run_ids"] = obj.PeerMaintenanceRunIds

	result["state"] = string(obj.LifecycleState)

	if obj.TargetDbServerVersion != nil {
		result["target_db_server_version"] = string(*obj.TargetDbServerVersion)
	}

	if obj.TargetResourceId != nil {
		result["target_resource_id"] = string(*obj.TargetResourceId)
	}

	result["target_resource_type"] = string(obj.TargetResourceType)

	if obj.TargetStorageServerVersion != nil {
		result["target_storage_server_version"] = string(*obj.TargetStorageServerVersion)
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

	if obj.TotalTimeTakenInMins != nil {
		result["total_time_taken_in_mins"] = int(*obj.TotalTimeTakenInMins)
	}

	return result
}
