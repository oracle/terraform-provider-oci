// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExecutionWindowDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["execution_window_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseExecutionWindowResource(), fieldMap, readSingularDatabaseExecutionWindow)
}

func readSingularDatabaseExecutionWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionWindowDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExecutionWindowDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExecutionWindowResponse
}

func (s *DatabaseExecutionWindowDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExecutionWindowDataSourceCrud) Get() error {
	request := oci_database.GetExecutionWindowRequest{}

	if executionWindowId, ok := s.D.GetOkExists("execution_window_id"); ok {
		tmp := executionWindowId.(string)
		request.ExecutionWindowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExecutionWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExecutionWindowDataSourceCrud) SetData() error {
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

	if s.Res.EstimatedTimeInMins != nil {
		s.D.Set("estimated_time_in_mins", *s.Res.EstimatedTimeInMins)
	}

	if s.Res.ExecutionResourceId != nil {
		s.D.Set("execution_resource_id", *s.Res.ExecutionResourceId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnforcedDuration != nil {
		s.D.Set("is_enforced_duration", *s.Res.IsEnforcedDuration)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
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

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalTimeTakenInMins != nil {
		s.D.Set("total_time_taken_in_mins", *s.Res.TotalTimeTakenInMins)
	}

	if s.Res.WindowDurationInMins != nil {
		s.D.Set("window_duration_in_mins", *s.Res.WindowDurationInMins)
	}

	s.D.Set("window_type", s.Res.WindowType)

	return nil
}
