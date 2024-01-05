// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseMaintenanceRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMaintenanceRuns,
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
			"maintenance_subtype": {
				Type:     schema.TypeString,
				Optional: true,
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
			"maintenance_runs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseMaintenanceRunResource()),
			},
		},
	}
}

func readDatabaseMaintenanceRuns(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMaintenanceRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListMaintenanceRunsResponse
}

func (s *DatabaseMaintenanceRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMaintenanceRunsDataSourceCrud) Get() error {
	request := oci_database.ListMaintenanceRunsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if maintenanceSubtype, ok := s.D.GetOkExists("maintenance_subtype"); ok {
		request.MaintenanceSubtype = oci_database.MaintenanceRunSummaryMaintenanceSubtypeEnum(maintenanceSubtype.(string))
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

	response, err := s.Client.ListMaintenanceRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaintenanceRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMaintenanceRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMaintenanceRunsDataSource-", DatabaseMaintenanceRunsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		maintenanceRun := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CurrentCustomActionTimeoutInMins != nil {
			maintenanceRun["current_custom_action_timeout_in_mins"] = *r.CurrentCustomActionTimeoutInMins
		}

		if r.CurrentPatchingComponent != nil {
			maintenanceRun["current_patching_component"] = *r.CurrentPatchingComponent
		}

		if r.CustomActionTimeoutInMins != nil {
			maintenanceRun["custom_action_timeout_in_mins"] = *r.CustomActionTimeoutInMins
		}

		if r.Description != nil {
			maintenanceRun["description"] = *r.Description
		}

		if r.DisplayName != nil {
			maintenanceRun["display_name"] = *r.DisplayName
		}

		if r.EstimatedComponentPatchingStartTime != nil {
			maintenanceRun["estimated_component_patching_start_time"] = r.EstimatedComponentPatchingStartTime.String()
		}

		if r.EstimatedPatchingTime != nil {
			maintenanceRun["estimated_patching_time"] = []interface{}{EstimatedPatchingTimeToMap(r.EstimatedPatchingTime)}
		} else {
			maintenanceRun["estimated_patching_time"] = nil
		}

		if r.Id != nil {
			maintenanceRun["id"] = *r.Id
		}

		if r.IsCustomActionTimeoutEnabled != nil {
			maintenanceRun["is_custom_action_timeout_enabled"] = *r.IsCustomActionTimeoutEnabled
		}

		if r.IsDstFileUpdateEnabled != nil {
			maintenanceRun["is_dst_file_update_enabled"] = *r.IsDstFileUpdateEnabled
		}

		if r.LifecycleDetails != nil {
			maintenanceRun["lifecycle_details"] = *r.LifecycleDetails
		}

		maintenanceRun["maintenance_subtype"] = r.MaintenanceSubtype

		maintenanceRun["maintenance_type"] = r.MaintenanceType

		if r.PatchFailureCount != nil {
			maintenanceRun["patch_failure_count"] = *r.PatchFailureCount
		}

		if r.PatchId != nil {
			maintenanceRun["patch_id"] = *r.PatchId
		}

		if r.PatchingEndTime != nil {
			maintenanceRun["patching_end_time"] = r.PatchingEndTime.String()
		}

		maintenanceRun["patching_mode"] = r.PatchingMode

		if r.PatchingStartTime != nil {
			maintenanceRun["patching_start_time"] = r.PatchingStartTime.String()
		}

		maintenanceRun["patching_status"] = r.PatchingStatus

		if r.PeerMaintenanceRunId != nil {
			maintenanceRun["peer_maintenance_run_id"] = *r.PeerMaintenanceRunId
		}

		maintenanceRun["state"] = r.LifecycleState

		if r.TargetDbServerVersion != nil {
			maintenanceRun["target_db_server_version"] = *r.TargetDbServerVersion
		}

		if r.TargetResourceId != nil {
			maintenanceRun["target_resource_id"] = *r.TargetResourceId
		}

		maintenanceRun["target_resource_type"] = r.TargetResourceType

		if r.TargetStorageServerVersion != nil {
			maintenanceRun["target_storage_server_version"] = *r.TargetStorageServerVersion
		}

		if r.TimeEnded != nil {
			maintenanceRun["time_ended"] = r.TimeEnded.String()
		}

		if r.TimeScheduled != nil {
			maintenanceRun["time_scheduled"] = r.TimeScheduled.Format(time.RFC3339Nano)
		}

		if r.TimeStarted != nil {
			maintenanceRun["time_started"] = r.TimeStarted.String()
		}

		resources = append(resources, maintenanceRun)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseMaintenanceRunsDataSource().Schema["maintenance_runs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("maintenance_runs", resources); err != nil {
		return err
	}

	return nil
}
