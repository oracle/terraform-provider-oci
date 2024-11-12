// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGatePipelineSchemasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGatePipelineSchemas,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pipeline_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pipeline_schema_collection": {
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
									"source_schema_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_schema_name": {
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

func readGoldenGatePipelineSchemas(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineSchemasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGatePipelineSchemasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListPipelineSchemasResponse
}

func (s *GoldenGatePipelineSchemasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGatePipelineSchemasDataSourceCrud) Get() error {
	request := oci_golden_gate.ListPipelineSchemasRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if pipelineId, ok := s.D.GetOkExists("pipeline_id"); ok {
		tmp := pipelineId.(string)
		request.PipelineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListPipelineSchemas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPipelineSchemas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGatePipelineSchemasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGatePipelineSchemasDataSource-", GoldenGatePipelineSchemasDataSource(), s.D))
	resources := []map[string]interface{}{}
	pipelineSchema := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PipelineSchemaSummaryToMap(item))
	}
	pipelineSchema["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGatePipelineSchemasDataSource().Schema["pipeline_schema_collection"].Elem.(*schema.Resource).Schema)
		pipelineSchema["items"] = items
	}

	resources = append(resources, pipelineSchema)
	if err := s.D.Set("pipeline_schema_collection", resources); err != nil {
		return err
	}

	return nil
}

func PipelineSchemaSummaryToMap(obj oci_golden_gate.PipelineSchemaSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SourceSchemaName != nil {
		result["source_schema_name"] = string(*obj.SourceSchemaName)
	}

	if obj.TargetSchemaName != nil {
		result["target_schema_name"] = string(*obj.TargetSchemaName)
	}

	return result
}
