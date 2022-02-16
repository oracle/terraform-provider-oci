// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v58/optimizer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OptimizerProfileDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["profile_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OptimizerProfileResource(), fieldMap, readSingularOptimizerProfile)
}

func readSingularOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.GetProfileResponse
}

func (s *OptimizerProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerProfileDataSourceCrud) Get() error {
	request := oci_optimizer.GetProfileRequest{}

	if profileId, ok := s.D.GetOkExists("profile_id"); ok {
		tmp := profileId.(string)
		request.ProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.GetProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OptimizerProfileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AggregationIntervalInDays != nil {
		s.D.Set("aggregation_interval_in_days", *s.Res.AggregationIntervalInDays)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LevelsConfiguration != nil {
		s.D.Set("levels_configuration", []interface{}{LevelsConfigurationToMap(s.Res.LevelsConfiguration)})
	} else {
		s.D.Set("levels_configuration", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetCompartments != nil {
		s.D.Set("target_compartments", []interface{}{TargetCompartmentsToMap(s.Res.TargetCompartments)})
	} else {
		s.D.Set("target_compartments", nil)
	}

	if s.Res.TargetTags != nil {
		s.D.Set("target_tags", []interface{}{TargetTagsToMap(s.Res.TargetTags)})
	} else {
		s.D.Set("target_tags", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
