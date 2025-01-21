// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceScheduleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["schedule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatascienceScheduleResource(), fieldMap, readSingularDatascienceSchedule)
}

func readSingularDatascienceSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceScheduleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceScheduleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetScheduleResponse
}

func (s *DatascienceScheduleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceScheduleDataSourceCrud) Get() error {
	request := oci_datascience.GetScheduleRequest{}

	if scheduleId, ok := s.D.GetOkExists("schedule_id"); ok {
		tmp := scheduleId.(string)
		request.ScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.GetSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceScheduleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Action != nil {
		actionArray := []interface{}{}
		if actionMap := ScheduleActionToMap(&s.Res.Action); actionMap != nil {
			actionArray = append(actionArray, actionMap)
		}
		s.D.Set("action", actionArray)
	} else {
		s.D.Set("action", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
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

	if s.Res.LastScheduleRunDetails != nil {
		s.D.Set("last_schedule_run_details", *s.Res.LastScheduleRunDetails)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogDetails != nil {
		s.D.Set("log_details", []interface{}{ScheduleLogDetailsToMap(s.Res.LogDetails)})
	} else {
		s.D.Set("log_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastScheduleRun != nil {
		s.D.Set("time_last_schedule_run", s.Res.TimeLastScheduleRun.String())
	}

	if s.Res.TimeNextScheduledRun != nil {
		s.D.Set("time_next_scheduled_run", s.Res.TimeNextScheduledRun.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Trigger != nil {
		triggerArray := []interface{}{}
		if triggerMap := ScheduleTriggerToMap(&s.Res.Trigger); triggerMap != nil {
			triggerArray = append(triggerArray, triggerMap)
		}
		s.D.Set("trigger", triggerArray)
	} else {
		s.D.Set("trigger", nil)
	}

	return nil
}
