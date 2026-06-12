// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiHostedApplicationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["hosted_application_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(GenerativeAiHostedApplicationResource(), fieldMap, readSingularGenerativeAiHostedApplicationWithContext)
}

func readSingularGenerativeAiHostedApplicationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedApplicationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiHostedApplicationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetHostedApplicationResponse
}

func (s *GenerativeAiHostedApplicationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiHostedApplicationDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetHostedApplicationRequest{}

	if hostedApplicationId, ok := s.D.GetOkExists("hosted_application_id"); ok {
		tmp := hostedApplicationId.(string)
		request.HostedApplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetHostedApplication(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiHostedApplicationDataSourceCrud) SetData() error {
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

	environmentVariables := []interface{}{}
	for _, item := range s.Res.EnvironmentVariables {
		environmentVariables = append(environmentVariables, EnvironmentVariableToMap(item))
	}
	s.D.Set("environment_variables", environmentVariables)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InboundAuthConfig != nil {
		s.D.Set("inbound_auth_config", []interface{}{InboundAuthConfigToMap(s.Res.InboundAuthConfig)})
	} else {
		s.D.Set("inbound_auth_config", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.NetworkingConfig != nil {
		s.D.Set("networking_config", []interface{}{NetworkingConfigToMap(s.Res.NetworkingConfig)})
	} else {
		s.D.Set("networking_config", nil)
	}

	if s.Res.ScalingConfig != nil {
		s.D.Set("scaling_config", []interface{}{ScalingConfigToMap(s.Res.ScalingConfig)})
	} else {
		s.D.Set("scaling_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	storageConfigs := []interface{}{}
	for _, item := range s.Res.StorageConfigs {
		storageConfigs = append(storageConfigs, StorageConfigToMap(item))
	}
	s.D.Set("storage_configs", storageConfigs)

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
