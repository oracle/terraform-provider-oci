// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceComputeTargetShapesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatascienceComputeTargetShapesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_target_shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"memory_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_series": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatascienceComputeTargetShapesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceComputeTargetShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatascienceComputeTargetShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListComputeTargetShapesResponse
}

func (s *DatascienceComputeTargetShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceComputeTargetShapesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datascience.ListComputeTargetShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListComputeTargetShapes(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeTargetShapes(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceComputeTargetShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceComputeTargetShapesDataSource-", DatascienceComputeTargetShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeTargetShape := map[string]interface{}{}

		if r.CoreCount != nil {
			computeTargetShape["core_count"] = *r.CoreCount
		}

		if r.MemoryInGBs != nil {
			computeTargetShape["memory_in_gbs"] = *r.MemoryInGBs
		}

		if r.Name != nil {
			computeTargetShape["name"] = *r.Name
		}

		computeTargetShape["shape_series"] = r.ShapeSeries

		resources = append(resources, computeTargetShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceComputeTargetShapesDataSource().Schema["compute_target_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compute_target_shapes", resources); err != nil {
		return err
	}

	return nil
}
