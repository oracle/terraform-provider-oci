// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseMaintenanceRunResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMaintenanceRun,
		Read:     readDatabaseMaintenanceRun,
		Update:   updateDatabaseMaintenanceRun,
		Delete:   deleteDatabaseMaintenanceRun,
		Schema: map[string]*schema.Schema{
			// Required
			"patch_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_scheduled": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},

			// Optional
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_dst_file_update_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"patching_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
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
			"is_custom_action_timeout_enabled": {
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
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseMaintenanceRunResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.MaintenanceRun
	DisableNotFoundRetries bool
}

type DatabaseMaintenanceUpdateResource struct {
	Details *oci_database.UpdateMaintenanceRunDetails
}

func (s *DatabaseMaintenanceRunResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseMaintenanceRunResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateScheduled),
		string(oci_database.MaintenanceRunLifecycleStateInProgress),
		string(oci_database.MaintenanceRunLifecycleStateUpdating),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateSucceeded),
		string(oci_database.MaintenanceRunLifecycleStateSkipped),
		string(oci_database.MaintenanceRunLifecycleStateScheduled),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateDeleting),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.MaintenanceRunLifecycleStateDeleted),
	}
}

func (s *DatabaseMaintenanceRunResourceCrud) Create() error {
	request := oci_database.CreateMaintenanceRunRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isDstFileUpdateEnabled, ok := s.D.GetOkExists("is_dst_file_update_enabled"); ok {
		tmp := isDstFileUpdateEnabled.(bool)
		request.IsDstFileUpdateEnabled = &tmp
	}

	if patchType, ok := s.D.GetOkExists("patch_type"); ok {
		request.PatchType = oci_database.CreateMaintenanceRunDetailsPatchTypeEnum(patchType.(string))
	}

	if patchingMode, ok := s.D.GetOkExists("patching_mode"); ok {
		request.PatchingMode = oci_database.CreateMaintenanceRunDetailsPatchingModeEnum(patchingMode.(string))
	}

	if targetResourceId, ok := s.D.GetOkExists("target_resource_id"); ok {
		tmp := targetResourceId.(string)
		request.TargetResourceId = &tmp
	}

	if timeScheduled, ok := s.D.GetOkExists("time_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeScheduled = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateMaintenanceRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceRun
	return nil
}

func (s *DatabaseMaintenanceRunResourceCrud) Get() error {
	request := oci_database.GetMaintenanceRunRequest{}

	tmp := s.D.Id()
	if tmp == "" {
		if id, ok := s.D.GetOkExists("maintenance_run_id"); ok {
			tmp = id.(string)
		}
	}
	request.MaintenanceRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetMaintenanceRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaintenanceRun
	return nil
}

func (s *DatabaseMaintenanceRunResourceCrud) Update() error {

	request := oci_database.UpdateMaintenanceRunRequest{}

	if currentCustomActionTimeoutInMins, ok := s.D.GetOkExists("current_custom_action_timeout_in_mins"); ok {
		tmp := currentCustomActionTimeoutInMins.(int)
		request.CurrentCustomActionTimeoutInMins = &tmp
	}

	if customActionTimeoutInMins, ok := s.D.GetOkExists("custom_action_timeout_in_mins"); ok {
		tmp := customActionTimeoutInMins.(int)
		request.CustomActionTimeoutInMins = &tmp
	}

	if isCustomActionTimeoutEnabled, ok := s.D.GetOkExists("is_custom_action_timeout_enabled"); ok {
		tmp := isCustomActionTimeoutEnabled.(bool)
		request.IsCustomActionTimeoutEnabled = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if isPatchNowEnabled, ok := s.D.GetOkExists("is_patch_now_enabled"); ok && s.D.HasChange("is_patch_now_enabled") {
		tmp := isPatchNowEnabled.(bool)
		request.IsPatchNowEnabled = &tmp
	}

	if isResumePatching, ok := s.D.GetOkExists("is_resume_patching"); ok {
		tmp := isResumePatching.(bool)
		request.IsResumePatching = &tmp
	}

	tmp := s.D.Id()
	request.MaintenanceRunId = &tmp

	if patchId, ok := s.D.GetOkExists("patch_id"); ok && s.D.HasChange("patchId") {
		tmp := patchId.(string)
		request.PatchId = &tmp
	}

	if patchingMode, ok := s.D.GetOkExists("patching_mode"); ok && s.D.HasChange("patching_mode") {
		request.PatchingMode = oci_database.UpdateMaintenanceRunDetailsPatchingModeEnum(patchingMode.(string))
	}

	if targetDbServerVersion, ok := s.D.GetOkExists("target_db_server_version"); ok {
		tmp := targetDbServerVersion.(string)
		request.TargetDbServerVersion = &tmp
	}

	if targetStorageServerVersion, ok := s.D.GetOkExists("target_storage_server_version"); ok {
		tmp := targetStorageServerVersion.(string)
		request.TargetStorageServerVersion = &tmp
	}

	if timeScheduled, ok := s.D.GetOkExists("time_scheduled"); ok {
		tmp, err := time.Parse(time.RFC3339, timeScheduled.(string))
		if err != nil {
			return err
		}
		request.TimeScheduled = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.UpdateMaintenanceRun(context.Background(), request)
	if err != nil {
		return err
	}
	// Workaround: Sleep for some time before polling the configuration. Because Update happens asynchronously, polling too
	// soon may result in service returning stale configuration values.
	time.Sleep(time.Second * 10)

	// Requests to Update may succeed instantly but may not see the actual Update take effect
	// until minutes later. Add polling here to return only when the change has taken effect.
	maintenanceRunUpdatePatchingModeFunc := func() bool { return s.Res.LifecycleState != oci_database.MaintenanceRunLifecycleStateUpdating }
	return tfresource.WaitForResourceCondition(s, maintenanceRunUpdatePatchingModeFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseMaintenanceRunResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentCustomActionTimeoutInMins != nil {
		s.D.Set("current_custom_action_timeout_in_mins", *s.Res.CurrentCustomActionTimeoutInMins)
	}

	if s.Res.CurrentPatchingComponent != nil {
		s.D.Set("current_patching_component", *s.Res.CurrentPatchingComponent)
	}

	if s.Res.CustomActionTimeoutInMins != nil {
		s.D.Set("custom_action_timeout_in_mins", *s.Res.CustomActionTimeoutInMins)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EstimatedComponentPatchingStartTime != nil {
		s.D.Set("estimated_component_patching_start_time", s.Res.EstimatedComponentPatchingStartTime.String())
	}

	if s.Res.EstimatedPatchingTime != nil {
		s.D.Set("estimated_patching_time", []interface{}{EstimatedPatchingTimeToMap(s.Res.EstimatedPatchingTime)})
	} else {
		s.D.Set("estimated_patching_time", nil)
	}

	if s.Res.IsCustomActionTimeoutEnabled != nil {
		s.D.Set("is_custom_action_timeout_enabled", *s.Res.IsCustomActionTimeoutEnabled)
	}

	if s.Res.IsDstFileUpdateEnabled != nil {
		s.D.Set("is_dst_file_update_enabled", *s.Res.IsDstFileUpdateEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_subtype", s.Res.MaintenanceSubtype)

	s.D.Set("maintenance_type", s.Res.MaintenanceType)

	if s.Res.PatchFailureCount != nil {
		s.D.Set("patch_failure_count", *s.Res.PatchFailureCount)
	}

	if s.Res.PatchId != nil {
		s.D.Set("patch_id", *s.Res.PatchId)
	}

	if s.Res.PatchingEndTime != nil {
		s.D.Set("patching_end_time", s.Res.PatchingEndTime.String())
	}

	s.D.Set("patching_mode", s.Res.PatchingMode)

	if s.Res.PatchingStartTime != nil {
		s.D.Set("patching_start_time", s.Res.PatchingStartTime.String())
	}

	s.D.Set("patching_status", s.Res.PatchingStatus)

	if s.Res.PeerMaintenanceRunId != nil {
		s.D.Set("peer_maintenance_run_id", *s.Res.PeerMaintenanceRunId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetDbServerVersion != nil {
		s.D.Set("target_db_server_version", *s.Res.TargetDbServerVersion)
	}

	if s.Res.TargetResourceId != nil {
		s.D.Set("target_resource_id", *s.Res.TargetResourceId)
	}

	s.D.Set("target_resource_type", s.Res.TargetResourceType)

	if s.Res.TargetStorageServerVersion != nil {
		s.D.Set("target_storage_server_version", *s.Res.TargetStorageServerVersion)
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeScheduled != nil {
		s.D.Set("time_scheduled", s.Res.TimeScheduled.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}

func EstimatedPatchingTimeToMap(obj *oci_database.EstimatedPatchingTime) map[string]interface{} {
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
