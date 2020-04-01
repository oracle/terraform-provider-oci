// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/dataflow"
)

func init() {
	RegisterDatasource("oci_dataflow_application", DataflowApplicationDataSource())
}

func DataflowApplicationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["application_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DataflowApplicationResource(), fieldMap, readSingularDataflowApplication)
}

func readSingularDataflowApplication(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowApplicationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataFlowClient

	return ReadResource(sync)
}

type DataflowApplicationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.GetApplicationResponse
}

func (s *DataflowApplicationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowApplicationDataSourceCrud) Get() error {
	request := oci_dataflow.GetApplicationRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "dataflow")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataflowApplicationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("arguments", s.Res.Arguments)

	if s.Res.ClassName != nil {
		s.D.Set("class_name", *s.Res.ClassName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("configuration", s.Res.Configuration)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DriverShape != nil {
		s.D.Set("driver_shape", *s.Res.DriverShape)
	}

	if s.Res.ExecutorShape != nil {
		s.D.Set("executor_shape", *s.Res.ExecutorShape)
	}

	if s.Res.FileUri != nil {
		s.D.Set("file_uri", *s.Res.FileUri)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("language", s.Res.Language)

	if s.Res.LogsBucketUri != nil {
		s.D.Set("logs_bucket_uri", *s.Res.LogsBucketUri)
	}

	if s.Res.NumExecutors != nil {
		s.D.Set("num_executors", *s.Res.NumExecutors)
	}

	if s.Res.OwnerPrincipalId != nil {
		s.D.Set("owner_principal_id", *s.Res.OwnerPrincipalId)
	}

	if s.Res.OwnerUserName != nil {
		s.D.Set("owner_user_name", *s.Res.OwnerUserName)
	}

	parameters := []interface{}{}
	for _, item := range s.Res.Parameters {
		parameters = append(parameters, ApplicationParameterToMap(item))
	}
	s.D.Set("parameters", parameters)

	if s.Res.SparkVersion != nil {
		s.D.Set("spark_version", *s.Res.SparkVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WarehouseBucketUri != nil {
		s.D.Set("warehouse_bucket_uri", *s.Res.WarehouseBucketUri)
	}

	return nil
}
