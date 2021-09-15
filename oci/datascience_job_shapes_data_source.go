// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v47/datascience"
)

func init() {
	RegisterDatasource("oci_datascience_job_shapes", DatascienceJobShapesDataSource())
}

func DatascienceJobShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceJobShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"job_shapes": {
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

func readDatascienceJobShapes(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return ReadResource(sync)
}

type DatascienceJobShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListJobShapesResponse
}

func (s *DatascienceJobShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceJobShapesDataSourceCrud) Get() error {
	request := oci_datascience.ListJobShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "datascience")

	response, err := s.Client.ListJobShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJobShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceJobShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatascienceJobShapesDataSource-", DatascienceJobShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		jobShape := map[string]interface{}{}

		if r.CoreCount != nil {
			jobShape["core_count"] = *r.CoreCount
		}

		if r.MemoryInGBs != nil {
			jobShape["memory_in_gbs"] = *r.MemoryInGBs
		}

		if r.Name != nil {
			jobShape["name"] = *r.Name
		}

		jobShape["shape_series"] = r.ShapeSeries

		resources = append(resources, jobShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatascienceJobShapesDataSource().Schema["job_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("job_shapes", resources); err != nil {
		return err
	}

	return nil
}
