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
			"memory_constraints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"max_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
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
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreImageShapeDataSource-", CoreImageShapeDataSource(), s.D))

	if s.Res.MemoryConstraints != nil {
		s.D.Set("memory_constraints", []interface{}{ImageMemoryConstraintsToMap(s.Res.MemoryConstraints)})
	} else {
		s.D.Set("memory_constraints", nil)
	}

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

func ImageMemoryConstraintsToMap(obj *oci_core.ImageMemoryConstraints) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxInGBs != nil {
		result["max_in_gbs"] = int(*obj.MaxInGBs)
	}

	if obj.MinInGBs != nil {
		result["min_in_gbs"] = int(*obj.MinInGBs)
	}

	return result
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
