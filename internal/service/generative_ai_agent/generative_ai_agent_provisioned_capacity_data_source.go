// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentProvisionedCapacityDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["provisioned_capacity_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(GenerativeAiAgentProvisionedCapacityResource(), fieldMap, readSingularGenerativeAiAgentProvisionedCapacityWithContext)
}

func readSingularGenerativeAiAgentProvisionedCapacityWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiAgentProvisionedCapacityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiAgentProvisionedCapacityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai_agent.GenerativeAiAgentClient
	Res    *oci_generative_ai_agent.GetProvisionedCapacityResponse
}

func (s *GenerativeAiAgentProvisionedCapacityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiAgentProvisionedCapacityDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai_agent.GetProvisionedCapacityRequest{}

	if provisionedCapacityId, ok := s.D.GetOkExists("provisioned_capacity_id"); ok {
		tmp := provisionedCapacityId.(string)
		request.ProvisionedCapacityId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai_agent")

	response, err := s.Client.GetProvisionedCapacity(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiAgentProvisionedCapacityDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.NumberOfUnits != nil {
		s.D.Set("number_of_units", *s.Res.NumberOfUnits)
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
