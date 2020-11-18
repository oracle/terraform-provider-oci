// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v29/core"
)

func init() {
	RegisterDatasource("oci_core_cpe_device_shapes", CoreCpeDeviceShapesDataSource())
}

func CoreCpeDeviceShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreCpeDeviceShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"cpe_device_shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cpe_device_info": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"platform_software_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vendor": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"cpe_device_shape_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"template": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreCpeDeviceShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCpeDeviceShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreCpeDeviceShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCpeDeviceShapesResponse
}

func (s *CoreCpeDeviceShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCpeDeviceShapesDataSourceCrud) Get() error {
	request := oci_core.ListCpeDeviceShapesRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListCpeDeviceShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCpeDeviceShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreCpeDeviceShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreCpeDeviceShapesDataSource-", CoreCpeDeviceShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cpeDeviceShape := map[string]interface{}{}

		if r.CpeDeviceInfo != nil {
			cpeDeviceShape["cpe_device_info"] = []interface{}{CpeDeviceInfoToMap(r.CpeDeviceInfo)}
		} else {
			cpeDeviceShape["cpe_device_info"] = nil
		}

		if r.Id != nil {
			cpeDeviceShape["cpe_device_shape_id"] = *r.Id
		}

		resources = append(resources, cpeDeviceShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreCpeDeviceShapesDataSource().Schema["cpe_device_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cpe_device_shapes", resources); err != nil {
		return err
	}

	return nil
}
