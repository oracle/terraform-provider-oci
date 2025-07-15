// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsTaskScheduleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fleet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: false,
		Optional: true,
	}
	fieldMap["task_schedule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsTaskScheduleResource(), fieldMap, readSingularJmsTaskSchedule)
}

func readSingularJmsTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &JmsTaskScheduleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsTaskScheduleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetTaskScheduleResponse
}

func (s *JmsTaskScheduleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsTaskScheduleDataSourceCrud) Get() error {
	request := oci_jms.GetTaskScheduleRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if taskScheduleId, ok := s.D.GetOkExists("task_schedule_id"); ok {
		tmp := taskScheduleId.(string)
		request.TaskScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsTaskScheduleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.ExecutionRecurrences != nil {
		s.D.Set("execution_recurrences", *s.Res.ExecutionRecurrences)
	}

	if s.Res.FleetId != nil {
		s.D.Set("fleet_id", *s.Res.FleetId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TaskDetails != nil {
		taskDetailsArray := []interface{}{}
		if taskDetailsMap := TaskDetailsToMap(&s.Res.TaskDetails); taskDetailsMap != nil {
			taskDetailsArray = append(taskDetailsArray, taskDetailsMap)
		}
		s.D.Set("task_details", taskDetailsArray)
	} else {
		s.D.Set("task_details", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRun != nil {
		s.D.Set("time_last_run", s.Res.TimeLastRun.String())
	}

	if s.Res.TimeLastUpdated != nil {
		s.D.Set("time_last_updated", s.Res.TimeLastUpdated.String())
	}

	if s.Res.TimeNextRun != nil {
		s.D.Set("time_next_run", s.Res.TimeNextRun.String())
	}

	return nil
}
