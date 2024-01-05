// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPsqlShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shape_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"memory_size_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ocpu_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"shape": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readPsqlShapes(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.ListShapesResponse
}

func (s *PsqlShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlShapesDataSourceCrud) Get() error {
	request := oci_psql.ListShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.ListShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *PsqlShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsqlShapesDataSource-", PsqlShapesDataSource(), s.D))
	resources := []map[string]interface{}{}
	shape := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ShapeSummaryToMap(item))
	}
	shape["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, PsqlShapesDataSource().Schema["shape_collection"].Elem.(*schema.Resource).Schema)
		shape["items"] = items
	}

	resources = append(resources, shape)
	if err := s.D.Set("shape_collection", resources); err != nil {
		return err
	}

	return nil
}

func ShapeSummaryToMap(obj oci_psql.ShapeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MemorySizeInGBs != nil {
		result["memory_size_in_gbs"] = int(*obj.MemorySizeInGBs)
	}

	if obj.OcpuCount != nil {
		result["ocpu_count"] = int(*obj.OcpuCount)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	return result
}
