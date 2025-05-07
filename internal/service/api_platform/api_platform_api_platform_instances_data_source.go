// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package api_platform

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_api_platform "github.com/oracle/oci-go-sdk/v65/apiplatform"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiPlatformApiPlatformInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApiPlatformApiPlatformInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"api_platform_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApiPlatformApiPlatformInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readApiPlatformApiPlatformInstances(d *schema.ResourceData, m interface{}) error {
	sync := &ApiPlatformApiPlatformInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApiPlatformClient()

	return tfresource.ReadResource(sync)
}

type ApiPlatformApiPlatformInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_api_platform.ApiPlatformClient
	Res    *oci_api_platform.ListApiPlatformInstancesResponse
}

func (s *ApiPlatformApiPlatformInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiPlatformApiPlatformInstancesDataSourceCrud) Get() error {
	request := oci_api_platform.ListApiPlatformInstancesRequest{}

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
		request.LifecycleState = oci_api_platform.ApiPlatformInstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "api_platform")

	response, err := s.Client.ListApiPlatformInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApiPlatformInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApiPlatformApiPlatformInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApiPlatformApiPlatformInstancesDataSource-", ApiPlatformApiPlatformInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	apiPlatformInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ApiPlatformInstanceSummaryToMap(item))
	}
	apiPlatformInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApiPlatformApiPlatformInstancesDataSource().Schema["api_platform_instance_collection"].Elem.(*schema.Resource).Schema)
		apiPlatformInstance["items"] = items
	}

	resources = append(resources, apiPlatformInstance)
	if err := s.D.Set("api_platform_instance_collection", resources); err != nil {
		return err
	}

	return nil
}
