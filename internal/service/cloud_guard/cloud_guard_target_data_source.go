// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v56/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CloudGuardTargetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["target_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardTargetResource(), fieldMap, readSingularCloudGuardTarget)
}

func readSingularCloudGuardTarget(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardTargetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

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

	response := *s.Res
	s.D.SetId(*response.GetId())

	if response.GetCompartmentId() != nil {
		s.D.Set("compartment_id", response.GetCompartmentId())
	}

	if s.Res.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.GetDefinedTags()))
	}

	if s.Res.GetDescription() != nil {
		s.D.Set("description", *s.Res.GetDescription())
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

	s.D.Set("inherited_by_compartments", s.Res.GetInheritedByCompartments())

	if s.Res.GetLifecyleDetails() != nil {
		s.D.Set("lifecyle_details", *s.Res.GetLifecyleDetails())
	}

	if s.Res.GetRecipeCount() != nil {
		s.D.Set("recipe_count", *s.Res.GetRecipeCount())
	}

	s.D.Set("state", s.Res.GetLifecycleState())

	if s.Res.GetSystemTags() != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.GetSystemTags()))
	}

	targetDetectorRecipes := []interface{}{}
	for _, item := range s.Res.GetTargetDetectorRecipes() {
		targetDetectorRecipes = append(targetDetectorRecipes, TargetDetectorRecipeToMap(item))
	}
	s.D.Set("target_detector_recipes", targetDetectorRecipes)

	if s.Res.GetTargetResourceId() != nil {
		s.D.Set("target_resource_id", *s.Res.GetTargetResourceId())
	}

	//s.D.Set("target_resource_type", s.Res.TargetResourceType())

	targetResponderRecipes := []interface{}{}
	for _, item := range s.Res.GetTargetResponderRecipes() {
		targetResponderRecipes = append(targetResponderRecipes, TargetResponderRecipeToMap(item))
	}
	s.D.Set("target_responder_recipes", targetResponderRecipes)

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.GetTimeUpdated().String())
	}

	return nil
}
