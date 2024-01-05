// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardDetectorRecipeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["detector_recipe_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardDetectorRecipeResource(), fieldMap, readSingularCloudGuardDetectorRecipe)
}

func readSingularCloudGuardDetectorRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDetectorRecipeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardDetectorRecipeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetDetectorRecipeResponse
}

func (s *CloudGuardDetectorRecipeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardDetectorRecipeDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetDetectorRecipeRequest{}

	if detectorRecipeId, ok := s.D.GetOkExists("detector_recipe_id"); ok {
		tmp := detectorRecipeId.(string)
		request.DetectorRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetDetectorRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardDetectorRecipeDataSourceCrud) SetData() error {
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

	s.D.Set("detector", s.Res.Detector)

	detectorRules := []interface{}{}
	for _, item := range s.Res.DetectorRules {
		detectorRules = append(detectorRules, DetectorRecipeDetectorRuleToMap(item))
	}
	s.D.Set("detector_rules", detectorRules)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	effectiveDetectorRules := []interface{}{}
	for _, item := range s.Res.EffectiveDetectorRules {
		effectiveDetectorRules = append(effectiveDetectorRules, DetectorRecipeDetectorRuleToMap(item))
	}
	s.D.Set("effective_detector_rules", effectiveDetectorRules)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("owner", s.Res.Owner)

	if s.Res.SourceDetectorRecipeId != nil {
		s.D.Set("source_detector_recipe_id", *s.Res.SourceDetectorRecipeId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("target_ids", s.Res.TargetIds)
	s.D.Set("target_ids", s.Res.TargetIds)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
