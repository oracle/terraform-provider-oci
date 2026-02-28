// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package gdp

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GdpGdpPipelineDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["gdp_pipeline_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["env"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  gdpCommercialCode,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(GdpGdpPipelineResource(), fieldMap, readSingularGdpGdpPipelineWithContext)
}

func readSingularGdpGdpPipelineWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GdpGdpPipelineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GuardedDataPipelineClient()
	if env, ok := sync.D.GetOk("env"); !ok || env.(string) != gdpUSGovCode {
		currentHost := sync.Client.Host
		newHost := strings.Replace(currentHost, "gdp", commercialSubdomain, 1)
		sync.Client.Host = newHost
	}

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GdpGdpPipelineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_gdp.GuardedDataPipelineClient
	Res    *oci_gdp.GetGdpPipelineResponse
}

func (s *GdpGdpPipelineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GdpGdpPipelineDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_gdp.GetGdpPipelineRequest{}

	if gdpPipelineId, ok := s.D.GetOkExists("gdp_pipeline_id"); ok {
		tmp := gdpPipelineId.(string)
		request.GdpPipelineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "gdp")

	response, err := s.Client.GetGdpPipeline(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GdpGdpPipelineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApprovalKeyVaultId != nil {
		s.D.Set("approval_key_vault_id", *s.Res.ApprovalKeyVaultId)
	}

	if s.Res.AuthorizationDetails != nil {
		s.D.Set("authorization_details", *s.Res.AuthorizationDetails)
	}

	bucketDetails := []interface{}{}
	for _, item := range s.Res.BucketDetails {
		bucketDetails = append(bucketDetails, BucketDetailsDefinitionToMap(item))
	}
	s.D.Set("bucket_details", bucketDetails)

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

	s.D.Set("file_types", s.Res.FileTypes)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsApprovalNeeded != nil {
		s.D.Set("is_approval_needed", *s.Res.IsApprovalNeeded)
	}

	if s.Res.IsChunkingEnabled != nil {
		s.D.Set("is_chunking_enabled", *s.Res.IsChunkingEnabled)
	}

	if s.Res.IsFileOverrideInDestinationEnabled != nil {
		s.D.Set("is_file_override_in_destination_enabled", *s.Res.IsFileOverrideInDestinationEnabled)
	}

	if s.Res.IsScanningEnabled != nil {
		s.D.Set("is_scanning_enabled", *s.Res.IsScanningEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeeredGdpPipelineId != nil {
		s.D.Set("peered_gdp_pipeline_id", *s.Res.PeeredGdpPipelineId)
	}

	if s.Res.PeeringRegion != nil {
		s.D.Set("peering_region", *s.Res.PeeringRegion)
	}

	s.D.Set("pipeline_type", s.Res.PipelineType)

	if s.Res.ServiceLogGroupId != nil {
		s.D.Set("service_log_group_id", *s.Res.ServiceLogGroupId)
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
