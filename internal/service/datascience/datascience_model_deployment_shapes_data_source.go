// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v58/datascience"
)

func DatascienceModelDeploymentShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModelDeploymentShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_deployment_shapes": {
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
					},
				},
			},
		},
	}
}

func readDatascienceModelDeploymentShapes(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelDeploymentShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelDeploymentShapesResponse
}

func (s *DatascienceModelDeploymentShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelDeploymentShapesDataSourceCrud) Get() error {
	request := oci_datascience.ListModelDeploymentShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelDeploymentShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelDeploymentShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelDeploymentShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelDeploymentShapesDataSource-", DatascienceModelDeploymentShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelDeploymentShape := map[string]interface{}{}

		if r.CoreCount != nil {
			modelDeploymentShape["core_count"] = *r.CoreCount
		}

		if r.MemoryInGBs != nil {
			modelDeploymentShape["memory_in_gbs"] = *r.MemoryInGBs
		}

		if r.Name != nil {
			modelDeploymentShape["name"] = *r.Name
		}

		resources = append(resources, modelDeploymentShape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelDeploymentShapesDataSource().Schema["model_deployment_shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_deployment_shapes", resources); err != nil {
		return err
	}

	return nil
}
