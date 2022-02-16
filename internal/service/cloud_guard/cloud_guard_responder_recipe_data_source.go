// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v58/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CloudGuardResponderRecipeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["responder_recipe_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudGuardResponderRecipeResource(), fieldMap, readSingularCloudGuardResponderRecipe)
}

func readSingularCloudGuardResponderRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResponderRecipeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardResponderRecipeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.GetResponderRecipeResponse
}

func (s *CloudGuardResponderRecipeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardResponderRecipeDataSourceCrud) Get() error {
	request := oci_cloud_guard.GetResponderRecipeRequest{}

	if responderRecipeId, ok := s.D.GetOkExists("responder_recipe_id"); ok {
		tmp := responderRecipeId.(string)
		request.ResponderRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.GetResponderRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardResponderRecipeDataSourceCrud) SetData() error {
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

	effectiveResponderRules := []interface{}{}
	for _, item := range s.Res.EffectiveResponderRules {
		effectiveResponderRules = append(effectiveResponderRules, ResponderRecipeResponderRuleToMap(item))
	}
	s.D.Set("effective_responder_rules", effectiveResponderRules)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("owner", s.Res.Owner)

	responderRules := []interface{}{}
	for _, item := range s.Res.ResponderRules {
		responderRules = append(responderRules, ResponderRecipeResponderRuleToMap(item))
	}
	s.D.Set("responder_rules", responderRules)

	if s.Res.SourceResponderRecipeId != nil {
		s.D.Set("source_responder_recipe_id", *s.Res.SourceResponderRecipeId)
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
