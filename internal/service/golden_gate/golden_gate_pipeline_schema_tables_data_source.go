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

func GoldenGatePipelineSchemaTablesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGatePipelineSchemaTables,
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
			"source_schema_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_schema_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pipeline_schema_table_collection": {
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
									"source_table_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_table_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
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
	}
}

func readGoldenGatePipelineSchemaTables(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineSchemaTablesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGatePipelineSchemaTablesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListPipelineSchemaTablesResponse
}

func (s *GoldenGatePipelineSchemaTablesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGatePipelineSchemaTablesDataSourceCrud) Get() error {
	request := oci_golden_gate.ListPipelineSchemaTablesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if pipelineId, ok := s.D.GetOkExists("pipeline_id"); ok {
		tmp := pipelineId.(string)
		request.PipelineId = &tmp
	}

	if sourceSchemaName, ok := s.D.GetOkExists("source_schema_name"); ok {
		tmp := sourceSchemaName.(string)
		request.SourceSchemaName = &tmp
	}

	if targetSchemaName, ok := s.D.GetOkExists("target_schema_name"); ok {
		tmp := targetSchemaName.(string)
		request.TargetSchemaName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListPipelineSchemaTables(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPipelineSchemaTables(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGatePipelineSchemaTablesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGatePipelineSchemaTablesDataSource-", GoldenGatePipelineSchemaTablesDataSource(), s.D))
	resources := []map[string]interface{}{}
	pipelineSchemaTable := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PipelineSchemaTableSummaryToMap(item))
	}
	pipelineSchemaTable["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGatePipelineSchemaTablesDataSource().Schema["pipeline_schema_table_collection"].Elem.(*schema.Resource).Schema)
		pipelineSchemaTable["items"] = items
	}

	resources = append(resources, pipelineSchemaTable)
	if err := s.D.Set("pipeline_schema_table_collection", resources); err != nil {
		return err
	}

	return nil
}

func PipelineSchemaTableSummaryToMap(obj oci_golden_gate.PipelineSchemaTableSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SourceTableName != nil {
		result["source_table_name"] = string(*obj.SourceTableName)
	}

	if obj.TargetTableName != nil {
		result["target_table_name"] = string(*obj.TargetTableName)
	}

	return result
}
