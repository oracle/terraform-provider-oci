// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeMaskingPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["masking_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeMaskingPolicyResource(), fieldMap, readSingularDataSafeMaskingPolicy)
}

func readSingularDataSafeMaskingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetMaskingPolicyResponse
}

func (s *DataSafeMaskingPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingPolicyDataSourceCrud) Get() error {
	request := oci_data_safe.GetMaskingPolicyRequest{}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetMaskingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeMaskingPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ColumnSource != nil {
		columnSourceArray := []interface{}{}
		if columnSourceMap := ColumnSourceDetailsToMap(&s.Res.ColumnSource); columnSourceMap != nil {
			columnSourceArray = append(columnSourceArray, columnSourceMap)
		}
		s.D.Set("column_source", columnSourceArray)
	} else {
		s.D.Set("column_source", nil)
	}

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

	if s.Res.IsDropTempTablesEnabled != nil {
		s.D.Set("is_drop_temp_tables_enabled", *s.Res.IsDropTempTablesEnabled)
	}

	if s.Res.IsRedoLoggingEnabled != nil {
		s.D.Set("is_redo_logging_enabled", *s.Res.IsRedoLoggingEnabled)
	}

	if s.Res.IsRefreshStatsEnabled != nil {
		s.D.Set("is_refresh_stats_enabled", *s.Res.IsRefreshStatsEnabled)
	}

	if s.Res.ParallelDegree != nil {
		s.D.Set("parallel_degree", *s.Res.ParallelDegree)
	}

	if s.Res.PostMaskingScript != nil {
		s.D.Set("post_masking_script", *s.Res.PostMaskingScript)
	}

	if s.Res.PreMaskingScript != nil {
		s.D.Set("pre_masking_script", *s.Res.PreMaskingScript)
	}

	s.D.Set("recompile", s.Res.Recompile)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
