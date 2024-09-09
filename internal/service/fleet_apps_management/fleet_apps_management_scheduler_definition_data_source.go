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

func FleetAppsManagementSchedulerDefinitionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["scheduler_definition_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementSchedulerDefinitionResource(), fieldMap, readSingularFleetAppsManagementSchedulerDefinition)
}

func readSingularFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerDefinitionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementSchedulerDefinitionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res    *oci_fleet_apps_management.GetSchedulerDefinitionResponse
}

func (s *FleetAppsManagementSchedulerDefinitionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementSchedulerDefinitionDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetSchedulerDefinitionRequest{}

	if schedulerDefinitionId, ok := s.D.GetOkExists("scheduler_definition_id"); ok {
		tmp := schedulerDefinitionId.(string)
		request.SchedulerDefinitionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetSchedulerDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementSchedulerDefinitionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("action_group_types", s.Res.ActionGroupTypes)

	actionGroups := []interface{}{}
	for _, item := range s.Res.ActionGroups {
		actionGroups = append(actionGroups, ActionGroupToMap(item))
	}
	s.D.Set("action_groups", actionGroups)

	if s.Res.ActivityInitiationCutOff != nil {
		s.D.Set("activity_initiation_cut_off", *s.Res.ActivityInitiationCutOff)
	}

	s.D.Set("application_types", s.Res.ApplicationTypes)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.CountOfAffectedActionGroups != nil {
		s.D.Set("count_of_affected_action_groups", *s.Res.CountOfAffectedActionGroups)
	}

	if s.Res.CountOfAffectedResources != nil {
		s.D.Set("count_of_affected_resources", *s.Res.CountOfAffectedResources)
	}

	if s.Res.CountOfAffectedTargets != nil {
		s.D.Set("count_of_affected_targets", *s.Res.CountOfAffectedTargets)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_operations", s.Res.LifecycleOperations)

	s.D.Set("products", s.Res.Products)

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	runBooks := []interface{}{}
	for _, item := range s.Res.RunBooks {
		runBooks = append(runBooks, OperationRunbookToMap(item))
	}
	s.D.Set("run_books", runBooks)

	if s.Res.Schedule != nil {
		s.D.Set("schedule", []interface{}{ScheduleToMap(s.Res.Schedule)})
	} else {
		s.D.Set("schedule", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfNextRun != nil {
		s.D.Set("time_of_next_run", s.Res.TimeOfNextRun.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
