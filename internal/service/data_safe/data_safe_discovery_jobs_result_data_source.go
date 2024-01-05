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

func DataSafeDiscoveryJobsResultDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["discovery_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["result_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeDiscoveryJobsResultResource(), fieldMap, readSingularDataSafeDiscoveryJobsResult)
}

func readSingularDataSafeDiscoveryJobsResult(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryJobsResultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeDiscoveryJobsResultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetDiscoveryJobResultResponse
}

func (s *DataSafeDiscoveryJobsResultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDiscoveryJobsResultDataSourceCrud) Get() error {
	request := oci_data_safe.GetDiscoveryJobResultRequest{}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	if resultKey, ok := s.D.GetOkExists("result_key"); ok {
		tmp := resultKey.(string)
		request.ResultKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetDiscoveryJobResult(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeDiscoveryJobsResultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeDiscoveryJobsResultDataSource-", DataSafeDiscoveryJobsResultDataSource(), s.D))

	s.D.Set("app_defined_child_column_keys", s.Res.AppDefinedChildColumnKeys)

	if s.Res.AppName != nil {
		s.D.Set("app_name", *s.Res.AppName)
	}

	if s.Res.ColumnName != nil {
		s.D.Set("column_name", *s.Res.ColumnName)
	}

	if s.Res.DataType != nil {
		s.D.Set("data_type", *s.Res.DataType)
	}

	s.D.Set("db_defined_child_column_keys", s.Res.DbDefinedChildColumnKeys)

	s.D.Set("discovery_type", s.Res.DiscoveryType)

	if s.Res.EstimatedDataValueCount != nil {
		s.D.Set("estimated_data_value_count", strconv.FormatInt(*s.Res.EstimatedDataValueCount, 10))
	}

	if s.Res.IsResultApplied != nil {
		s.D.Set("is_result_applied", *s.Res.IsResultApplied)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.ModifiedAttributes != nil {
		s.D.Set("modified_attributes", []interface{}{ModifiedAttributesToMap(s.Res.ModifiedAttributes)})
	} else {
		s.D.Set("modified_attributes", nil)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("object_type", s.Res.ObjectType)

	s.D.Set("parent_column_keys", s.Res.ParentColumnKeys)

	s.D.Set("planned_action", s.Res.PlannedAction)

	s.D.Set("relation_type", s.Res.RelationType)

	s.D.Set("sample_data_values", s.Res.SampleDataValues)

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.SensitiveColumnkey != nil {
		s.D.Set("sensitive_columnkey", *s.Res.SensitiveColumnkey)
	}

	if s.Res.SensitiveTypeId != nil {
		s.D.Set("sensitive_type_id", *s.Res.SensitiveTypeId)
	}

	return nil
}
