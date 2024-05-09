// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubScheduledJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["scheduled_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubScheduledJobResource(), fieldMap, readSingularOsManagementHubScheduledJob)
}

func readSingularOsManagementHubScheduledJob(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubScheduledJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledJobClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubScheduledJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.ScheduledJobClient
	Res    *oci_os_management_hub.GetScheduledJobResponse
}

func (s *OsManagementHubScheduledJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubScheduledJobDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetScheduledJobRequest{}

	if scheduledJobId, ok := s.D.GetOkExists("scheduled_job_id"); ok {
		tmp := scheduledJobId.(string)
		request.ScheduledJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetScheduledJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubScheduledJobDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsManagedByAutonomousLinux != nil {
		s.D.Set("is_managed_by_autonomous_linux", *s.Res.IsManagedByAutonomousLinux)
	}

	if s.Res.IsRestricted != nil {
		s.D.Set("is_restricted", *s.Res.IsRestricted)
	}

	if s.Res.IsSubcompartmentIncluded != nil {
		s.D.Set("is_subcompartment_included", *s.Res.IsSubcompartmentIncluded)
	}

	s.D.Set("lifecycle_stage_ids", s.Res.LifecycleStageIds)

	s.D.Set("locations", s.Res.Locations)

	s.D.Set("managed_compartment_ids", s.Res.ManagedCompartmentIds)

	s.D.Set("managed_instance_group_ids", s.Res.ManagedInstanceGroupIds)

	s.D.Set("managed_instance_ids", s.Res.ManagedInstanceIds)

	operations := []interface{}{}
	for _, item := range s.Res.Operations {
		operations = append(operations, ScheduledJobOperationToMap(item))
	}
	s.D.Set("operations", operations)

	if s.Res.RecurringRule != nil {
		s.D.Set("recurring_rule", *s.Res.RecurringRule)
	}

	s.D.Set("retry_intervals", s.Res.RetryIntervals)

	s.D.Set("schedule_type", s.Res.ScheduleType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastExecution != nil {
		s.D.Set("time_last_execution", s.Res.TimeLastExecution.String())
	}

	if s.Res.TimeNextExecution != nil {
		s.D.Set("time_next_execution", s.Res.TimeNextExecution.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("work_request_ids", s.Res.WorkRequestIds)

	return nil
}
