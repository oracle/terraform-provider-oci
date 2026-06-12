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

func GenerativeAiHostedDeploymentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["hosted_deployment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(GenerativeAiHostedDeploymentResource(), fieldMap, readSingularGenerativeAiHostedDeploymentWithContext)
}

func readSingularGenerativeAiHostedDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiHostedDeploymentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiHostedDeploymentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetHostedDeploymentResponse
}

func (s *GenerativeAiHostedDeploymentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiHostedDeploymentDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetHostedDeploymentRequest{}

	if hostedDeploymentId, ok := s.D.GetOkExists("hosted_deployment_id"); ok {
		tmp := hostedDeploymentId.(string)
		request.HostedDeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetHostedDeployment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiHostedDeploymentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActiveArtifact != nil {
		activeArtifactArray := []interface{}{}
		if activeArtifactMap := ArtifactToMap(s.Res.ActiveArtifact); activeArtifactMap != nil {
			activeArtifactArray = append(activeArtifactArray, activeArtifactMap)
		}
		s.D.Set("active_artifact", activeArtifactArray)
	} else {
		s.D.Set("active_artifact", nil)
	}

	artifacts := []interface{}{}
	for _, item := range s.Res.Artifacts {
		artifacts = append(artifacts, ArtifactToMap(item))
	}
	s.D.Set("artifacts", artifacts)

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

	if s.Res.HostedApplicationId != nil {
		s.D.Set("hosted_application_id", *s.Res.HostedApplicationId)
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
