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

func DatabaseSchedulingPlanDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["scheduling_plan_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseSchedulingPlanResource(), fieldMap, readSingularDatabaseSchedulingPlan)
}

func readSingularDatabaseSchedulingPlan(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPlanDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSchedulingPlanDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetSchedulingPlanResponse
}

func (s *DatabaseSchedulingPlanDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSchedulingPlanDataSourceCrud) Get() error {
	request := oci_database.GetSchedulingPlanRequest{}

	if schedulingPlanId, ok := s.D.GetOkExists("scheduling_plan_id"); ok {
		tmp := schedulingPlanId.(string)
		request.SchedulingPlanId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetSchedulingPlan(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseSchedulingPlanDataSourceCrud) SetData() error {
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

	if s.Res.EstimatedTimeInMins != nil {
		s.D.Set("estimated_time_in_mins", *s.Res.EstimatedTimeInMins)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsUsingRecommendedScheduledActions != nil {
		s.D.Set("is_using_recommended_scheduled_actions", *s.Res.IsUsingRecommendedScheduledActions)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("plan_intent", s.Res.PlanIntent)

	if s.Res.ResourceId != nil {
		s.D.Set("resource_id", *s.Res.ResourceId)
	}

	if s.Res.SchedulingPolicyId != nil {
		s.D.Set("scheduling_policy_id", *s.Res.SchedulingPolicyId)
	}

	s.D.Set("service_type", s.Res.ServiceType)

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
