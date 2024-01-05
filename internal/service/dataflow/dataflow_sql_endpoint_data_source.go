// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataflowSqlEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sql_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataflowSqlEndpointResource(), fieldMap, readSingularDataflowSqlEndpoint)
}

func readSingularDataflowSqlEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowSqlEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowSqlEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.GetSqlEndpointResponse
}

func (s *DataflowSqlEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowSqlEndpointDataSourceCrud) Get() error {
	request := oci_dataflow.GetSqlEndpointRequest{}

	if sqlEndpointId, ok := s.D.GetOkExists("sql_endpoint_id"); ok {
		tmp := sqlEndpointId.(string)
		request.SqlEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.GetSqlEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataflowSqlEndpointDataSourceCrud) SetData() error {
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

	if s.Res.DriverShape != nil {
		s.D.Set("driver_shape", *s.Res.DriverShape)
	}

	if s.Res.DriverShapeConfig != nil {
		s.D.Set("driver_shape_config", []interface{}{ShapeConfigToMap(s.Res.DriverShapeConfig)})
	} else {
		s.D.Set("driver_shape_config", nil)
	}

	if s.Res.ExecutorShape != nil {
		s.D.Set("executor_shape", *s.Res.ExecutorShape)
	}

	if s.Res.ExecutorShapeConfig != nil {
		s.D.Set("executor_shape_config", []interface{}{ShapeConfigToMap(s.Res.ExecutorShapeConfig)})
	} else {
		s.D.Set("executor_shape_config", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.JdbcEndpointUrl != nil {
		s.D.Set("jdbc_endpoint_url", *s.Res.JdbcEndpointUrl)
	}

	if s.Res.LakeId != nil {
		s.D.Set("lake_id", *s.Res.LakeId)
	}

	if s.Res.LastAcceptedRequestToken != nil {
		s.D.Set("last_accepted_request_token", *s.Res.LastAcceptedRequestToken)
	}

	if s.Res.MaxExecutorCount != nil {
		s.D.Set("max_executor_count", *s.Res.MaxExecutorCount)
	}

	if s.Res.MetastoreId != nil {
		s.D.Set("metastore_id", *s.Res.MetastoreId)
	}

	if s.Res.MinExecutorCount != nil {
		s.D.Set("min_executor_count", *s.Res.MinExecutorCount)
	}

	if s.Res.NetworkConfiguration != nil {
		networkConfigurationArray := []interface{}{}
		if networkConfigurationMap := SqlEndpointNetworkConfigurationToMap(&s.Res.NetworkConfiguration, true); networkConfigurationMap != nil {
			networkConfigurationArray = append(networkConfigurationArray, networkConfigurationMap)
		}
		s.D.Set("network_configuration", networkConfigurationArray)
	} else {
		s.D.Set("network_configuration", nil)
	}

	s.D.Set("spark_advanced_configurations", s.Res.SparkAdvancedConfigurations)

	if s.Res.SqlEndpointVersion != nil {
		s.D.Set("sql_endpoint_version", *s.Res.SqlEndpointVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

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
