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

func DataflowPoolsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataflowPools,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_starts_with": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"owner_principal_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DataflowPoolResource()),
						},
					},
				},
			},
		},
	}
}

func readDataflowPools(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowPoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowPoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.ListPoolsResponse
}

func (s *DataflowPoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowPoolsDataSourceCrud) Get() error {
	request := oci_dataflow.ListPoolsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameStartsWith, ok := s.D.GetOkExists("display_name_starts_with"); ok {
		tmp := displayNameStartsWith.(string)
		request.DisplayNameStartsWith = &tmp
	}

	if ownerPrincipalId, ok := s.D.GetOkExists("owner_principal_id"); ok {
		tmp := ownerPrincipalId.(string)
		request.OwnerPrincipalId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dataflow.ListPoolsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.ListPools(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPools(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataflowPoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataflowPoolsDataSource-", DataflowPoolsDataSource(), s.D))
	resources := []map[string]interface{}{}
	pool := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PoolSummaryToMap(item))
	}
	pool["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataflowPoolsDataSource().Schema["pool_collection"].Elem.(*schema.Resource).Schema)
		pool["items"] = items
	}

	resources = append(resources, pool)
	if err := s.D.Set("pool_collection", resources); err != nil {
		return err
	}

	return nil
}
