// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v56/opsi"
)

func OpsiOperationsInsightsWarehousesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiOperationsInsightsWarehouses,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"operations_insights_warehouse_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpsiOperationsInsightsWarehouseResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiOperationsInsightsWarehouses(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehousesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOperationsInsightsWarehousesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListOperationsInsightsWarehousesResponse
}

func (s *OpsiOperationsInsightsWarehousesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOperationsInsightsWarehousesDataSourceCrud) Get() error {
	request := oci_opsi.ListOperationsInsightsWarehousesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		//		request.LifecycleState = oci_opsi.ListOperationsInsightsWarehousesLifecycleStateEnum(state.(string))
		interfaces := state.([]interface{})
		tmp := make([]oci_opsi.OperationsInsightsWarehouseLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.OperationsInsightsWarehouseLifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListOperationsInsightsWarehouses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOperationsInsightsWarehouses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiOperationsInsightsWarehousesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiOperationsInsightsWarehousesDataSource-", OpsiOperationsInsightsWarehousesDataSource(), s.D))
	resources := []map[string]interface{}{}
	operationsInsightsWarehouse := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OperationsInsightsWarehouseSummaryToMap(item))
	}
	operationsInsightsWarehouse["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiOperationsInsightsWarehousesDataSource().Schema["operations_insights_warehouse_summary_collection"].Elem.(*schema.Resource).Schema)
		operationsInsightsWarehouse["items"] = items
	}

	resources = append(resources, operationsInsightsWarehouse)
	if err := s.D.Set("operations_insights_warehouse_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
