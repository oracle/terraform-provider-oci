// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreVirtualCircuitBandwidthShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVirtualCircuitBandwidthShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"provider_service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_circuit_bandwidth_shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"bandwidth_in_mbps": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreVirtualCircuitBandwidthShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVirtualCircuitBandwidthShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreVirtualCircuitBandwidthShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListFastConnectProviderVirtualCircuitBandwidthShapesResponse
}

func (s *CoreVirtualCircuitBandwidthShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVirtualCircuitBandwidthShapesDataSourceCrud) Get() error {
	request := oci_core.ListFastConnectProviderVirtualCircuitBandwidthShapesRequest{}

	if providerServiceId, ok := s.D.GetOkExists("provider_service_id"); ok {
		tmp := providerServiceId.(string)
		request.ProviderServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListFastConnectProviderVirtualCircuitBandwidthShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFastConnectProviderVirtualCircuitBandwidthShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVirtualCircuitBandwidthShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVirtualCircuitBandwidthShapesDataSource-", CoreVirtualCircuitBandwidthShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		virtualCircuitBandwidthShape := map[string]interface{}{}

		if r.BandwidthInMbps != nil {
			virtualCircuitBandwidthShape["bandwidth_in_mbps"] = *r.BandwidthInMbps
		}

		if r.Name != nil {
			virtualCircuitBandwidthShape["name"] = *r.Name
		}

		resources = append(resources, virtualCircuitBandwidthShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVirtualCircuitBandwidthShapesDataSource().Schema["virtual_circuit_bandwidth_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("virtual_circuit_bandwidth_shapes", resources); err != nil {
		return err
	}

	return nil
}
