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

func CoreImageShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreImageShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_shape_compatibilities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"image_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_constraints": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
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
							MaxItems: 1,
							MinItems: 1,
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
				},
			},
		},
	}
}

func readCoreImageShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreImageShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListImageShapeCompatibilityEntriesResponse
}

func (s *CoreImageShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreImageShapesDataSourceCrud) Get() error {
	request := oci_core.ListImageShapeCompatibilityEntriesRequest{}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListImageShapeCompatibilityEntries(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListImageShapeCompatibilityEntries(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreImageShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreImageShapesDataSource-", CoreImageShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		imageShape := map[string]interface{}{
			"image_id": *r.ImageId,
		}

		if r.MemoryConstraints != nil {
			imageShape["memory_constraints"] = []interface{}{ImageMemoryConstraintsToMap(r.MemoryConstraints)}
		} else {
			imageShape["memory_constraints"] = nil
		}

		if r.OcpuConstraints != nil {
			imageShape["ocpu_constraints"] = []interface{}{ImageOcpuConstraintsToMap(r.OcpuConstraints)}
		} else {
			imageShape["ocpu_constraints"] = nil
		}

		if r.Shape != nil {
			imageShape["shape"] = *r.Shape
		}

		resources = append(resources, imageShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreImageShapesDataSource().Schema["image_shape_compatibilities"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("image_shape_compatibilities", resources); err != nil {
		return err
	}

	return nil
}
