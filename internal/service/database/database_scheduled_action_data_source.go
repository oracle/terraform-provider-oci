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

func DatabaseScheduledActionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["scheduled_action_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseScheduledActionResource(), fieldMap, readSingularDatabaseScheduledAction)
}

func readSingularDatabaseScheduledAction(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseScheduledActionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseScheduledActionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetScheduledActionResponse
}

func (s *DatabaseScheduledActionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseScheduledActionDataSourceCrud) Get() error {
	request := oci_database.GetScheduledActionRequest{}

	if scheduledActionId, ok := s.D.GetOkExists("scheduled_action_id"); ok {
		tmp := scheduledActionId.(string)
		request.ScheduledActionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetScheduledAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseScheduledActionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	actionMembers := []interface{}{}
	for _, item := range s.Res.ActionMembers {
		actionMembers = append(actionMembers, ActionMemberToMap(item))
	}
	s.D.Set("action_members", actionMembers)

	if s.Res.ActionOrder != nil {
		s.D.Set("action_order", *s.Res.ActionOrder)
	}

	s.D.Set("action_params", s.Res.ActionParams)

	s.D.Set("action_type", s.Res.ActionType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EstimatedTimeInMins != nil {
		s.D.Set("estimated_time_in_mins", *s.Res.EstimatedTimeInMins)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SchedulingPlanId != nil {
		s.D.Set("scheduling_plan_id", *s.Res.SchedulingPlanId)
	}

	if s.Res.SchedulingWindowId != nil {
		s.D.Set("scheduling_window_id", *s.Res.SchedulingWindowId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
