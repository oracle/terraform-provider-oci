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

func ServiceMeshAccessPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceMeshAccessPolicies,
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
			"mesh_id": {
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
			"access_policy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ServiceMeshAccessPolicyResource()),
						},
					},
				},
			},
		},
	}
}

func readServiceMeshAccessPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceMeshAccessPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceMeshClient()

	return tfresource.ReadResource(sync)
}

type ServiceMeshAccessPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_mesh.ServiceMeshClient
	Res    *oci_service_mesh.ListAccessPoliciesResponse
}

func (s *ServiceMeshAccessPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceMeshAccessPoliciesDataSourceCrud) Get() error {
	request := oci_service_mesh.ListAccessPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if meshId, ok := s.D.GetOkExists("mesh_id"); ok {
		tmp := meshId.(string)
		request.MeshId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_service_mesh.AccessPolicyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_mesh")

	response, err := s.Client.ListAccessPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAccessPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceMeshAccessPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceMeshAccessPoliciesDataSource-", ServiceMeshAccessPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}
	accessPolicy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AccessPolicySummaryToMap(item))
	}
	accessPolicy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ServiceMeshAccessPoliciesDataSource().Schema["access_policy_collection"].Elem.(*schema.Resource).Schema)
		accessPolicy["items"] = items
	}

	resources = append(resources, accessPolicy)
	if err := s.D.Set("access_policy_collection", resources); err != nil {
		return err
	}

	return nil
}
