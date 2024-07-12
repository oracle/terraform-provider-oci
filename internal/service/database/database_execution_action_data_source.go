// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExecutionActionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["execution_action_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseExecutionActionResource(), fieldMap, readSingularDatabaseExecutionAction)
}

func readSingularDatabaseExecutionAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExecutionActionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExecutionActionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExecutionActionResponse
}

func (s *DatabaseExecutionActionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExecutionActionDataSourceCrud) Get() error {
	request := oci_database.GetExecutionActionRequest{}

	if executionActionId, ok := s.D.GetOkExists("execution_action_id"); ok {
		tmp := executionActionId.(string)
		request.ExecutionActionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExecutionAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExecutionActionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	actionMembers := []interface{}{}
	for _, item := range s.Res.ActionMembers {
		actionMembers = append(actionMembers, ExecutionActionMemberToMap(item))
	}
	s.D.Set("action_members", actionMembers)

	s.D.Set("action_params", s.Res.ActionParams)

	s.D.Set("action_type", s.Res.ActionType)

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

	if s.Res.ExecutionActionOrder != nil {
		s.D.Set("execution_action_order", *s.Res.ExecutionActionOrder)
	}

	if s.Res.ExecutionWindowId != nil {
		s.D.Set("execution_window_id", *s.Res.ExecutionWindowId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalTimeTakenInMins != nil {
		s.D.Set("total_time_taken_in_mins", *s.Res.TotalTimeTakenInMins)
	}

	return nil
}
