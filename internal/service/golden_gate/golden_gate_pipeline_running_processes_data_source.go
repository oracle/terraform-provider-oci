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

func GoldenGatePipelineRunningProcessesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGatePipelineRunningProcesses,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"pipeline_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pipeline_running_process_collection": {
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
									"last_record_lag_in_seconds": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"process_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_processed": {
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

func readGoldenGatePipelineRunningProcesses(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineRunningProcessesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGatePipelineRunningProcessesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListPipelineRunningProcessesResponse
}

func (s *GoldenGatePipelineRunningProcessesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGatePipelineRunningProcessesDataSourceCrud) Get() error {
	request := oci_golden_gate.ListPipelineRunningProcessesRequest{}

	if pipelineId, ok := s.D.GetOkExists("pipeline_id"); ok {
		tmp := pipelineId.(string)
		request.PipelineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListPipelineRunningProcesses(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPipelineRunningProcesses(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGatePipelineRunningProcessesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGatePipelineRunningProcessesDataSource-", GoldenGatePipelineRunningProcessesDataSource(), s.D))
	resources := []map[string]interface{}{}
	pipelineRunningProcess := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PipelineRunningProcessSummaryToMap(item))
	}
	pipelineRunningProcess["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGatePipelineRunningProcessesDataSource().Schema["pipeline_running_process_collection"].Elem.(*schema.Resource).Schema)
		pipelineRunningProcess["items"] = items
	}

	resources = append(resources, pipelineRunningProcess)
	if err := s.D.Set("pipeline_running_process_collection", resources); err != nil {
		return err
	}

	return nil
}

func PipelineRunningProcessSummaryToMap(obj oci_golden_gate.PipelineRunningProcessSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LastRecordLagInSeconds != nil {
		result["last_record_lag_in_seconds"] = float32(*obj.LastRecordLagInSeconds)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["process_type"] = string(obj.ProcessType)

	result["status"] = string(obj.Status)

	if obj.TimeLastProcessed != nil {
		result["time_last_processed"] = obj.TimeLastProcessed.String()
	}

	return result
}
