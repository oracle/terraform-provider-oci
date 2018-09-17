// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CrossConnectPortSpeedShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCrossConnectPortSpeedShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cross_connect_port_speed_shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_speed_in_gbps": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCrossConnectPortSpeedShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectPortSpeedShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type CrossConnectPortSpeedShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCrossconnectPortSpeedShapesResponse
}

func (s *CrossConnectPortSpeedShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CrossConnectPortSpeedShapesDataSourceCrud) Get() error {
	request := oci_core.ListCrossconnectPortSpeedShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListCrossconnectPortSpeedShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCrossconnectPortSpeedShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CrossConnectPortSpeedShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		crossConnectPortSpeedShape := map[string]interface{}{}

		if r.Name != nil {
			crossConnectPortSpeedShape["name"] = *r.Name
		}

		if r.PortSpeedInGbps != nil {
			crossConnectPortSpeedShape["port_speed_in_gbps"] = *r.PortSpeedInGbps
		}

		resources = append(resources, crossConnectPortSpeedShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CrossConnectPortSpeedShapesDataSource().Schema["cross_connect_port_speed_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cross_connect_port_speed_shapes", resources); err != nil {
		return err
	}

	return nil
}
