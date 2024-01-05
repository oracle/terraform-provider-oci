// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSensitiveDataModelsSensitiveColumnDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sensitive_column_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["sensitive_data_model_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeSensitiveDataModelsSensitiveColumnResource(), fieldMap, readSingularDataSafeSensitiveDataModelsSensitiveColumn)
}

func readSingularDataSafeSensitiveDataModelsSensitiveColumn(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSensitiveDataModelsSensitiveColumnDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSensitiveDataModelsSensitiveColumnDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSensitiveColumnResponse
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnDataSourceCrud) Get() error {
	request := oci_data_safe.GetSensitiveColumnRequest{}

	if sensitiveColumnKey, ok := s.D.GetOkExists("sensitive_column_key"); ok {
		tmp := sensitiveColumnKey.(string)
		request.SensitiveColumnKey = &tmp
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSensitiveColumn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSensitiveDataModelsSensitiveColumnDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSensitiveDataModelsSensitiveColumnDataSource-", DataSafeSensitiveDataModelsSensitiveColumnDataSource(), s.D))

	if s.Res.SensitiveDataModelId != nil {
		s.D.Set("sensitive_data_model_id", *s.Res.SensitiveDataModelId)
	}

	s.D.Set("app_defined_child_column_keys", s.Res.AppDefinedChildColumnKeys)

	if s.Res.AppName != nil {
		s.D.Set("app_name", *s.Res.AppName)
	}

	s.D.Set("column_groups", s.Res.ColumnGroups)

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	if s.Res.DataType != nil {
		s.D.Set("data_type", *s.Res.DataType)
	}

	s.D.Set("db_defined_child_column_keys", s.Res.DbDefinedChildColumnKeys)

	if s.Res.EstimatedDataValueCount != nil {
		s.D.Set("estimated_data_value_count", strconv.FormatInt(*s.Res.EstimatedDataValueCount, 10))
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("object_type", s.Res.ObjectType)

	s.D.Set("parent_column_keys", s.Res.ParentColumnKeys)

	s.D.Set("relation_type", s.Res.RelationType)

	s.D.Set("sample_data_values", s.Res.SampleDataValues)

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.SensitiveTypeId != nil {
		s.D.Set("sensitive_type_id", *s.Res.SensitiveTypeId)
	}

	s.D.Set("source", s.Res.Source)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
