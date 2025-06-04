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

func FleetAppsManagementRunbookVersionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["runbook_version_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementRunbookVersionResource(), fieldMap, readSingularFleetAppsManagementRunbookVersion)
}

func readSingularFleetAppsManagementRunbookVersion(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementRunbookVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	Res    *oci_fleet_apps_management.GetRunbookVersionResponse
}

func (s *FleetAppsManagementRunbookVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementRunbookVersionDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetRunbookVersionRequest{}

	if runbookVersionId, ok := s.D.GetOkExists("runbook_version_id"); ok {
		tmp := runbookVersionId.(string)
		request.RunbookVersionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetRunbookVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementRunbookVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.ExecutionWorkflowDetails != nil {
		s.D.Set("execution_workflow_details", []interface{}{ExecutionWorkflowDetailsToMap(s.Res.ExecutionWorkflowDetails)})
	} else {
		s.D.Set("execution_workflow_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	groups := []interface{}{}
	for _, item := range s.Res.Groups {
		groups = append(groups, GroupToMap(item))
	}
	s.D.Set("groups", groups)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.RollbackWorkflowDetails != nil {
		s.D.Set("rollback_workflow_details", []interface{}{RollbackWorkflowDetailsToMap(s.Res.RollbackWorkflowDetails)})
	} else {
		s.D.Set("rollback_workflow_details", nil)
	}

	if s.Res.RunbookId != nil {
		s.D.Set("runbook_id", *s.Res.RunbookId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	tasks := []interface{}{}
	for _, item := range s.Res.Tasks {
		tasks = append(tasks, TaskToMap(item))
	}
	s.D.Set("tasks", tasks)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
