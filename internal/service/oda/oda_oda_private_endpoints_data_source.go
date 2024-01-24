// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OdaOdaPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOdaOdaPrivateEndpoints,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oda_private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OdaOdaPrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readOdaOdaPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.ManagementClient
	Res    *oci_oda.ListOdaPrivateEndpointsResponse
}

func (s *OdaOdaPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_oda.ListOdaPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_oda.OdaPrivateEndpointLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.ListOdaPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOdaPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OdaOdaPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OdaOdaPrivateEndpointsDataSource-", OdaOdaPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	odaPrivateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OdaPrivateEndpointSummaryToMap(item))
	}
	odaPrivateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OdaOdaPrivateEndpointsDataSource().Schema["oda_private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		odaPrivateEndpoint["items"] = items
	}

	resources = append(resources, odaPrivateEndpoint)
	if err := s.D.Set("oda_private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
