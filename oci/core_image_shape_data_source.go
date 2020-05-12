// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func init() {
	RegisterDatasource("oci_core_image_shape", CoreImageShapeDataSource())
}

func CoreImageShapeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreImageShape,
		Schema: map[string]*schema.Schema{
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"ocpu_constraints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"max": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreImageShape(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageShapeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreImageShapeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetImageShapeCompatibilityEntryResponse
}

func (s *CoreImageShapeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreImageShapeDataSourceCrud) Get() error {
	request := oci_core.GetImageShapeCompatibilityEntryRequest{}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if shapeName, ok := s.D.GetOkExists("shape_name"); ok {
		tmp := shapeName.(string)
		request.ShapeName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetImageShapeCompatibilityEntry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreImageShapeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	if s.Res.OcpuConstraints != nil {
		s.D.Set("ocpu_constraints", []interface{}{ImageOcpuConstraintsToMap(s.Res.OcpuConstraints)})
	} else {
		s.D.Set("ocpu_constraints", nil)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	return nil
}

func ImageOcpuConstraintsToMap(obj *oci_core.ImageOcpuConstraints) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = int(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	return result
}
