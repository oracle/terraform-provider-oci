// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_data_platform

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_data_platform "github.com/oracle/oci-go-sdk/v65/aidataplatform"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiDataPlatformAiDataPlatformDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ai_data_platform_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(AiDataPlatformAiDataPlatformResource(), fieldMap, readSingularAiDataPlatformAiDataPlatformWithContext)
}

func readSingularAiDataPlatformAiDataPlatformWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AiDataPlatformAiDataPlatformDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiDataPlatformClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type AiDataPlatformAiDataPlatformDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ai_data_platform.AiDataPlatformClient
	Res    *oci_ai_data_platform.GetAiDataPlatformResponse
}

func (s *AiDataPlatformAiDataPlatformDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AiDataPlatformAiDataPlatformDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ai_data_platform.GetAiDataPlatformRequest{}

	if aiDataPlatformId, ok := s.D.GetOkExists("ai_data_platform_id"); ok {
		tmp := aiDataPlatformId.(string)
		request.AiDataPlatformId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ai_data_platform")

	response, err := s.Client.GetAiDataPlatform(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AiDataPlatformAiDataPlatformDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AiDataPlatformType != nil {
		s.D.Set("ai_data_platform_type", *s.Res.AiDataPlatformType)
	}

	if s.Res.AliasKey != nil {
		s.D.Set("alias_key", *s.Res.AliasKey)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
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

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WebSocketEndpoint != nil {
		s.D.Set("web_socket_endpoint", *s.Res.WebSocketEndpoint)
	}

	return nil
}
