// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_scheduler

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resource_scheduler "github.com/oracle/oci-go-sdk/v65/resourcescheduler"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceSchedulerScheduleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["schedule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ResourceSchedulerScheduleResource(), fieldMap, readSingularResourceSchedulerSchedule)
}

func readSingularResourceSchedulerSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceSchedulerScheduleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduleClient()

	return tfresource.ReadResource(sync)
}

type ResourceSchedulerScheduleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_scheduler.ScheduleClient
	Res    *oci_resource_scheduler.GetScheduleResponse
}

func (s *ResourceSchedulerScheduleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceSchedulerScheduleDataSourceCrud) Get() error {
	request := oci_resource_scheduler.GetScheduleRequest{}

	if scheduleId, ok := s.D.GetOkExists("schedule_id"); ok {
		tmp := scheduleId.(string)
		request.ScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_scheduler")

	response, err := s.Client.GetSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourceSchedulerScheduleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("action", s.Res.Action)

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

	if s.Res.RecurrenceDetails != nil {
		s.D.Set("recurrence_details", *s.Res.RecurrenceDetails)
	}

	s.D.Set("recurrence_type", s.Res.RecurrenceType)

	resourceFilters := []interface{}{}
	for _, item := range s.Res.ResourceFilters {
		resourceFilters = append(resourceFilters, ResourceFilterToMap(item))
	}
	s.D.Set("resource_filters", resourceFilters)

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, ResourceToMap(item))
	}
	s.D.Set("resources", resources)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnds != nil {
		s.D.Set("time_ends", s.Res.TimeEnds.Format(time.RFC3339Nano))
	}

	if s.Res.TimeLastRun != nil {
		s.D.Set("time_last_run", s.Res.TimeLastRun.String())
	}

	if s.Res.TimeNextRun != nil {
		s.D.Set("time_next_run", s.Res.TimeNextRun.String())
	}

	if s.Res.TimeStarts != nil {
		s.D.Set("time_starts", s.Res.TimeStarts.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
