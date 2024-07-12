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

func DatabaseSchedulingPolicySchedulingWindowDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["scheduling_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["scheduling_window_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseSchedulingPolicySchedulingWindowResource(), fieldMap, readSingularDatabaseSchedulingPolicySchedulingWindow)
}

func readSingularDatabaseSchedulingPolicySchedulingWindow(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicySchedulingWindowDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSchedulingPolicySchedulingWindowDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetSchedulingWindowResponse
}

func (s *DatabaseSchedulingPolicySchedulingWindowDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSchedulingPolicySchedulingWindowDataSourceCrud) Get() error {
	request := oci_database.GetSchedulingWindowRequest{}

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	if schedulingWindowId, ok := s.D.GetOkExists("scheduling_window_id"); ok {
		tmp := schedulingWindowId.(string)
		request.SchedulingWindowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetSchedulingWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseSchedulingPolicySchedulingWindowDataSourceCrud) SetData() error {
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeNextSchedulingWindowStarts != nil {
		s.D.Set("time_next_scheduling_window_starts", s.Res.TimeNextSchedulingWindowStarts.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WindowPreference != nil {
		s.D.Set("window_preference", []interface{}{WindowPreferenceDetailToMap(s.Res.WindowPreference)})
	} else {
		s.D.Set("window_preference", nil)
	}

	return nil
}
