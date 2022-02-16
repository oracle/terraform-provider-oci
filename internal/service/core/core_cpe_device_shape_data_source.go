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

func CoreCpeDeviceShapeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreCpeDeviceShape,
		Schema: map[string]*schema.Schema{
			"cpe_device_shape_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"cpe_device_info": {
				Type:     schema.TypeList,
				Computed: true,
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
			"parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"explanation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"template": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreCpeDeviceShape(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCpeDeviceShapeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreCpeDeviceShapeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCpeDeviceShapeResponse
}

func (s *CoreCpeDeviceShapeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCpeDeviceShapeDataSourceCrud) Get() error {
	request := oci_core.GetCpeDeviceShapeRequest{}

	if cpeDeviceShapeId, ok := s.D.GetOkExists("cpe_device_shape_id"); ok {
		tmp := cpeDeviceShapeId.(string)
		request.CpeDeviceShapeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetCpeDeviceShape(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreCpeDeviceShapeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.CpeDeviceShapeId)

	if s.Res.CpeDeviceInfo != nil {
		s.D.Set("cpe_device_info", []interface{}{CpeDeviceInfoToMap(s.Res.CpeDeviceInfo)})
	} else {
		s.D.Set("cpe_device_info", nil)
	}

	parameters := []interface{}{}
	for _, item := range s.Res.Parameters {
		parameters = append(parameters, CpeDeviceConfigQuestionToMap(item))
	}
	s.D.Set("parameters", parameters)

	if s.Res.Template != nil {
		s.D.Set("template", *s.Res.Template)
	}

	return nil
}

func CpeDeviceConfigQuestionToMap(obj oci_core.CpeDeviceConfigQuestion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Explanation != nil {
		result["explanation"] = string(*obj.Explanation)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	return result
}

func CpeDeviceInfoToMap(obj *oci_core.CpeDeviceInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PlatformSoftwareVersion != nil {
		result["platform_software_version"] = string(*obj.PlatformSoftwareVersion)
	}

	if obj.Vendor != nil {
		result["vendor"] = string(*obj.Vendor)
	}

	return result
}
