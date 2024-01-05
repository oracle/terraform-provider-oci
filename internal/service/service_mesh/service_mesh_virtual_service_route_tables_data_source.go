// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_mesh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceMeshVirtualServiceRouteTablesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceMeshVirtualServiceRouteTables,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service_route_table_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ServiceMeshVirtualServiceRouteTableResource()),
						},
					},
				},
			},
		},
	}
}

func readServiceMeshVirtualServiceRouteTables(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualServiceRouteTablesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshVirtualServiceRouteTablesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.ListVirtualServiceRouteTablesResponse
}

func (s *ServiceMeshVirtualServiceRouteTablesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshVirtualServiceRouteTablesDataSourceCrud) Get() error {
	request := oci_service_mesh.ListVirtualServiceRouteTablesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_service_mesh.VirtualServiceRouteTableLifecycleStateEnum(state.(string))
	}

	if virtualServiceId, ok := s.D.GetOkExists("virtual_service_id"); ok {
		tmp := virtualServiceId.(string)
		request.VirtualServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.ListVirtualServiceRouteTables(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVirtualServiceRouteTables(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceMeshVirtualServiceRouteTablesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceMeshVirtualServiceRouteTablesDataSource-", ServiceMeshVirtualServiceRouteTablesDataSource(), s.D))
	resources := []map[string]interface{}{}
	virtualServiceRouteTable := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, VirtualServiceRouteTableSummaryToMap(item))
	}
	virtualServiceRouteTable["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ServiceMeshVirtualServiceRouteTablesDataSource().Schema["virtual_service_route_table_collection"].Elem.(*schema.Resource).Schema)
		virtualServiceRouteTable["items"] = items
	}

	resources = append(resources, virtualServiceRouteTable)
	if err := s.D.Set("virtual_service_route_table_collection", resources); err != nil {
		return err
	}

	return nil
}
