// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v25/cloudguard"
)

func init() {
	RegisterDatasource("oci_cloud_guard_target", CloudGuardTargetDataSource())
}

func CloudGuardTargetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["target_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CloudGuardTargetResource(), fieldMap, readSingularCloudGuardTarget)
}

func readSingularCloudGuardTarget(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardTargetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).cloudGuardClient()

	return ReadResource(sync)
}

type CloudGuardTargetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetTargetResponse
}

func (s *CloudGuardTargetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardTargetDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetTargetRequest{}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardTargetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("inherited_by_compartments", s.Res.InheritedByCompartments)

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.RecipeCount != nil {
		s.D.Set("recipe_count", *s.Res.RecipeCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	targetDetectorRecipes := []interface{}{}
	for _, item := range s.Res.TargetDetectorRecipes {
		targetDetectorRecipes = append(targetDetectorRecipes, TargetDetectorRecipeToMap(item))
	}
	s.D.Set("target_detector_recipes", targetDetectorRecipes)

	if s.Res.TargetResourceId != nil {
		s.D.Set("target_resource_id", *s.Res.TargetResourceId)
	}

	s.D.Set("target_resource_type", s.Res.TargetResourceType)

	targetResponderRecipes := []interface{}{}
	for _, item := range s.Res.TargetResponderRecipes {
		targetResponderRecipes = append(targetResponderRecipes, TargetResponderRecipeToMap(item))
	}
	s.D.Set("target_responder_recipes", targetResponderRecipes)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
