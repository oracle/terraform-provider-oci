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

func DataflowSqlEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataflowSqlEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sql_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataflowSqlEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readDataflowSqlEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowSqlEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowSqlEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.ListSqlEndpointsResponse
}

func (s *DataflowSqlEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowSqlEndpointsDataSourceCrud) Get() error {
	request := oci_dataflow.ListSqlEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sqlEndpointId, ok := s.D.GetOkExists("id"); ok {
		tmp := sqlEndpointId.(string)
		request.SqlEndpointId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dataflow.ListSqlEndpointsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.ListSqlEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSqlEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataflowSqlEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataflowSqlEndpointsDataSource-", DataflowSqlEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	sqlEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SqlEndpointSummaryToMap(item))
	}
	sqlEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataflowSqlEndpointsDataSource().Schema["sql_endpoint_collection"].Elem.(*schema.Resource).Schema)
		sqlEndpoint["items"] = items
	}

	resources = append(resources, sqlEndpoint)
	if err := s.D.Set("sql_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
