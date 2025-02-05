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

func DatabaseSchedulingPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["scheduling_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseSchedulingPolicyResource(), fieldMap, readSingularDatabaseSchedulingPolicy)
}

func readSingularDatabaseSchedulingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseSchedulingPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseSchedulingPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetSchedulingPolicyResponse
}

func (s *DatabaseSchedulingPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseSchedulingPolicyDataSourceCrud) Get() error {
	request := oci_database.GetSchedulingPolicyRequest{}

	if schedulingPolicyId, ok := s.D.GetOkExists("scheduling_policy_id"); ok {
		tmp := schedulingPolicyId.(string)
		request.SchedulingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetSchedulingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseSchedulingPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("cadence", s.Res.Cadence)

	if s.Res.CadenceStartMonth != nil {
		s.D.Set("cadence_start_month", []interface{}{MonthToMapPolicy(s.Res.CadenceStartMonth)})
	} else {
		s.D.Set("cadence_start_month", nil)
	}

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

	if s.Res.TimeNextWindowStarts != nil {
		s.D.Set("time_next_window_starts", s.Res.TimeNextWindowStarts.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
