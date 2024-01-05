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

func OdaOdaPrivateEndpointScanProxiesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOdaOdaPrivateEndpointScanProxies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"oda_private_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oda_private_endpoint_scan_proxy_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OdaOdaPrivateEndpointScanProxyResource()),
						},
					},
				},
			},
		},
	}
}

func readOdaOdaPrivateEndpointScanProxies(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointScanProxiesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaPrivateEndpointScanProxiesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.ManagementClient
	Res    *oci_oda.ListOdaPrivateEndpointScanProxiesResponse
}

func (s *OdaOdaPrivateEndpointScanProxiesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaPrivateEndpointScanProxiesDataSourceCrud) Get() error {
	request := oci_oda.ListOdaPrivateEndpointScanProxiesRequest{}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_oda.OdaPrivateEndpointScanProxyLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.ListOdaPrivateEndpointScanProxies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOdaPrivateEndpointScanProxies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OdaOdaPrivateEndpointScanProxiesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OdaOdaPrivateEndpointScanProxiesDataSource-", OdaOdaPrivateEndpointScanProxiesDataSource(), s.D))
	resources := []map[string]interface{}{}
	odaPrivateEndpointScanProxy := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OdaPrivateEndpointScanProxySummaryToMap(item))
	}
	odaPrivateEndpointScanProxy["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OdaOdaPrivateEndpointScanProxiesDataSource().Schema["oda_private_endpoint_scan_proxy_collection"].Elem.(*schema.Resource).Schema)
		odaPrivateEndpointScanProxy["items"] = items
	}

	resources = append(resources, odaPrivateEndpointScanProxy)
	if err := s.D.Set("oda_private_endpoint_scan_proxy_collection", resources); err != nil {
		return err
	}

	return nil
}
