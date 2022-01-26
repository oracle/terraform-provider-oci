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

func OpsiOperationsInsightsWarehouseUsersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiOperationsInsightsWarehouseUsers,
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
			"operations_insights_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"operations_insights_warehouse_user_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpsiOperationsInsightsWarehouseUserResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiOperationsInsightsWarehouseUsers(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseUsersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOperationsInsightsWarehouseUsersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListOperationsInsightsWarehouseUsersResponse
}

func (s *OpsiOperationsInsightsWarehouseUsersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOperationsInsightsWarehouseUsersDataSourceCrud) Get() error {
	request := oci_opsi.ListOperationsInsightsWarehouseUsersRequest{}

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

	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		request.OperationsInsightsWarehouseId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		//		request.LifecycleState = oci_opsi.ListOperationsInsightsWarehouseUsersLifecycleStateEnum(state.(string))
		interfaces := state.([]interface{})
		tmp := make([]oci_opsi.OperationsInsightsWarehouseUserLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.OperationsInsightsWarehouseUserLifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListOperationsInsightsWarehouseUsers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOperationsInsightsWarehouseUsers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiOperationsInsightsWarehouseUsersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiOperationsInsightsWarehouseUsersDataSource-", OpsiOperationsInsightsWarehouseUsersDataSource(), s.D))
	resources := []map[string]interface{}{}
	operationsInsightsWarehouseUser := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OperationsInsightsWarehouseUserSummaryToMap(item))
	}
	operationsInsightsWarehouseUser["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiOperationsInsightsWarehouseUsersDataSource().Schema["operations_insights_warehouse_user_summary_collection"].Elem.(*schema.Resource).Schema)
		operationsInsightsWarehouseUser["items"] = items
	}

	resources = append(resources, operationsInsightsWarehouseUser)
	if err := s.D.Set("operations_insights_warehouse_user_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
