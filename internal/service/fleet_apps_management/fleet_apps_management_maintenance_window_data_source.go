// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementMaintenanceWindowDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["maintenance_window_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementMaintenanceWindowResource(), fieldMap, readSingularFleetAppsManagementMaintenanceWindow)
}

func readSingularFleetAppsManagementMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementMaintenanceWindowDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementMaintenanceWindowClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementMaintenanceWindowDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementMaintenanceWindowClient
	Res    *oci_fleet_apps_management.GetMaintenanceWindowResponse
}

func (s *FleetAppsManagementMaintenanceWindowDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementMaintenanceWindowDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetMaintenanceWindowRequest{}

	if maintenanceWindowId, ok := s.D.GetOkExists("maintenance_window_id"); ok {
		tmp := maintenanceWindowId.(string)
		request.MaintenanceWindowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementMaintenanceWindowDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Duration != nil {
		s.D.Set("duration", *s.Res.Duration)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOutage != nil {
		s.D.Set("is_outage", *s.Res.IsOutage)
	}

	if s.Res.IsRecurring != nil {
		s.D.Set("is_recurring", *s.Res.IsRecurring)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_window_type", s.Res.MaintenanceWindowType)

	if s.Res.Recurrences != nil {
		s.D.Set("recurrences", *s.Res.Recurrences)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TaskInitiationCutoff != nil {
		s.D.Set("task_initiation_cutoff", *s.Res.TaskInitiationCutoff)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeScheduleStart != nil {
		s.D.Set("time_schedule_start", s.Res.TimeScheduleStart.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
