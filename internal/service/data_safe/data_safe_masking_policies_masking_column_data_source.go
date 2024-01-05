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

func DataSafeMaskingPoliciesMaskingColumnDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["masking_column_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["masking_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeMaskingPoliciesMaskingColumnResource(), fieldMap, readSingularDataSafeMaskingPoliciesMaskingColumn)
}

func readSingularDataSafeMaskingPoliciesMaskingColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPoliciesMaskingColumnDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingPoliciesMaskingColumnDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetMaskingColumnResponse
}

func (s *DataSafeMaskingPoliciesMaskingColumnDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingPoliciesMaskingColumnDataSourceCrud) Get() error {
	request := oci_data_safe.GetMaskingColumnRequest{}

	if maskingColumnKey, ok := s.D.GetOkExists("masking_column_key"); ok {
		tmp := maskingColumnKey.(string)
		request.MaskingColumnKey = &tmp
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetMaskingColumn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeMaskingPoliciesMaskingColumnDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingPoliciesMaskingColumnDataSource-", DataSafeMaskingPoliciesMaskingColumnDataSource(), s.D))

	s.D.Set("child_columns", s.Res.ChildColumns)

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	if s.Res.DataType != nil {
		s.D.Set("data_type", *s.Res.DataType)
	}

	if s.Res.IsMaskingEnabled != nil {
		s.D.Set("is_masking_enabled", *s.Res.IsMaskingEnabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaskingColumnGroup != nil {
		s.D.Set("masking_column_group", *s.Res.MaskingColumnGroup)
	}

	maskingFormats := []interface{}{}
	for _, item := range s.Res.MaskingFormats {
		maskingFormats = append(maskingFormats, MaskingFormatToMap(item))
	}
	s.D.Set("masking_formats", maskingFormats)

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("object_type", s.Res.ObjectType)

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.SensitiveTypeId != nil {
		s.D.Set("sensitive_type_id", *s.Res.SensitiveTypeId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
