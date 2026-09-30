// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subsetting_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DataSafeSubsettingPolicyResource(), fieldMap, readSingularDataSafeSubsettingPolicyWithContext)
}

func readSingularDataSafeSubsettingPolicyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSubsettingPolicyResponse
}

func (s *DataSafeSubsettingPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicyDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingPolicyRequest{}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSubsettingPolicy(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSubsettingPolicyDataSourceCrud) SetData() error {
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

	if s.Res.IsRedoLoggingEnabled != nil {
		s.D.Set("is_redo_logging_enabled", *s.Res.IsRedoLoggingEnabled)
	}

	if s.Res.IsRefreshStatsEnabled != nil {
		s.D.Set("is_refresh_stats_enabled", *s.Res.IsRefreshStatsEnabled)
	}

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.ParallelDegree != nil {
		s.D.Set("parallel_degree", *s.Res.ParallelDegree)
	}

	if s.Res.PostSubsettingScript != nil {
		s.D.Set("post_subsetting_script", *s.Res.PostSubsettingScript)
	}

	if s.Res.PreSubsettingScript != nil {
		s.D.Set("pre_subsetting_script", *s.Res.PreSubsettingScript)
	}

	s.D.Set("recompile", s.Res.Recompile)

	if s.Res.SchemaSource != nil {
		schemaSourceArray := []interface{}{}
		if schemaSourceMap := SchemaSourceDetailsToMap(&s.Res.SchemaSource); schemaSourceMap != nil {
			schemaSourceArray = append(schemaSourceArray, schemaSourceMap)
		}
		s.D.Set("schema_source", schemaSourceArray)
	} else {
		s.D.Set("schema_source", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("unrelated_tables_action", s.Res.UnrelatedTablesAction)

	return nil
}
