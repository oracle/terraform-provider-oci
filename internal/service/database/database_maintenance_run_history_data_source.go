// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMaintenanceRunHistoryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseMaintenanceRunHistory,
		Schema: map[string]*schema.Schema{
			"maintenance_run_history_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
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
					},
				},
			},
		},
	}
}

func readSingularDatabaseMaintenanceRunHistory(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunHistoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMaintenanceRunHistoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetMaintenanceRunHistoryResponse
}

func (s *DatabaseMaintenanceRunHistoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMaintenanceRunHistoryDataSourceCrud) Get() error {
	request := oci_database.GetMaintenanceRunHistoryRequest{}

	if maintenanceRunHistoryId, ok := s.D.GetOkExists("maintenance_run_history_id"); ok {
		tmp := maintenanceRunHistoryId.(string)
		request.MaintenanceRunHistoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetMaintenanceRunHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMaintenanceRunHistoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	dbServersHistoryDetails := []interface{}{}
	for _, item := range s.Res.DbServersHistoryDetails {
		dbServersHistoryDetails = append(dbServersHistoryDetails, DbServerHstrySummaryToMap(item))
	}
	s.D.Set("db_servers_history_details", dbServersHistoryDetails)

	if s.Res.MaintenanceRunDetails != nil {
		s.D.Set("maintenance_run_details", []interface{}{MtnanceRunSummaryToMap(s.Res.MaintenanceRunDetails)})
	} else {
		s.D.Set("maintenance_run_details", nil)
	}

	return nil
}

func DbServerHstrySummaryToMap(obj oci_database.DbServerHistorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbServerPatchingDetails != nil {
		result["db_server_patching_details"] = []interface{}{DbServerPatchingDtlsToMap(obj.DbServerPatchingDetails)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func DbServerPatchingDtlsToMap(obj *oci_database.DbServerPatchingDetails) map[string]interface{} {
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

func EstmdPatchingTimeToMap(obj *oci_database.EstimatedPatchingTime) map[string]interface{} {
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

func MtnanceRunSummaryToMap(obj *oci_database.MaintenanceRunSummary) map[string]interface{} {
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
		result["estimated_patching_time"] = []interface{}{EstmdPatchingTimeToMap(obj.EstimatedPatchingTime)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCustomActionTimeoutEnabled != nil {
		result["is_custom_action_timeout_enabled"] = bool(*obj.IsCustomActionTimeoutEnabled)
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

	return result
}
