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

func ServiceMeshVirtualDeploymentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceMeshVirtualDeployments,
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
			"virtual_deployment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ServiceMeshVirtualDeploymentResource()),
						},
					},
				},
			},
		},
	}
}

func readServiceMeshVirtualDeployments(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshVirtualDeploymentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshVirtualDeploymentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.ListVirtualDeploymentsResponse
}

func (s *ServiceMeshVirtualDeploymentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshVirtualDeploymentsDataSourceCrud) Get() error {
	request := oci_service_mesh.ListVirtualDeploymentsRequest{}

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
		request.LifecycleState = oci_service_mesh.VirtualDeploymentLifecycleStateEnum(state.(string))
	}

	if virtualServiceId, ok := s.D.GetOkExists("virtual_service_id"); ok {
		tmp := virtualServiceId.(string)
		request.VirtualServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.ListVirtualDeployments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVirtualDeployments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceMeshVirtualDeploymentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceMeshVirtualDeploymentsDataSource-", ServiceMeshVirtualDeploymentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	virtualDeployment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, VirtualDeploymentSummaryToMap(item))
	}
	virtualDeployment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ServiceMeshVirtualDeploymentsDataSource().Schema["virtual_deployment_collection"].Elem.(*schema.Resource).Schema)
		virtualDeployment["items"] = items
	}

	resources = append(resources, virtualDeployment)
	if err := s.D.Set("virtual_deployment_collection", resources); err != nil {
		return err
	}

	return nil
}
