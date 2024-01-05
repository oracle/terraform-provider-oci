// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcemanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v65/resourcemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourcemanagerPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readResourcemanagerPrivateEndpoints,
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
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ResourcemanagerPrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readResourcemanagerPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &ResourcemanagerPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceManagerClient()

	return tfresource.ReadResource(sync)
}

type ResourcemanagerPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resourcemanager.ResourceManagerClient
	Res    *oci_resourcemanager.ListPrivateEndpointsResponse
}

func (s *ResourcemanagerPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourcemanagerPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_resourcemanager.ListPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if privateEndpointId, ok := s.D.GetOkExists("id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resourcemanager")

	response, err := s.Client.ListPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ResourcemanagerPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ResourcemanagerPrivateEndpointsDataSource-", ResourcemanagerPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	privateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PrivateEndpointSummaryToMap(item))
	}
	privateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ResourcemanagerPrivateEndpointsDataSource().Schema["private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		privateEndpoint["items"] = items
	}

	resources = append(resources, privateEndpoint)
	if err := s.D.Set("private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
