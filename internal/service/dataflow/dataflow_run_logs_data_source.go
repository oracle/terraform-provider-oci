// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v58/dataflow"
)

func DataflowRunLogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataflowRunLogs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"run_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"run_logs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"run_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_in_bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataflowRunLogs(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowRunLogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowRunLogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.ListRunLogsResponse
}

func (s *DataflowRunLogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowRunLogsDataSourceCrud) Get() error {
	request := oci_dataflow.ListRunLogsRequest{}

	if runId, ok := s.D.GetOkExists("run_id"); ok {
		tmp := runId.(string)
		request.RunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.ListRunLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRunLogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataflowRunLogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataflowRunLogsDataSource-", DataflowRunLogsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		runLog := map[string]interface{}{
			"run_id": *r.RunId,
		}

		if r.Name != nil {
			runLog["name"] = *r.Name
		}

		if r.SizeInBytes != nil {
			runLog["size_in_bytes"] = strconv.FormatInt(*r.SizeInBytes, 10)
		}

		runLog["source"] = r.Source

		if r.TimeCreated != nil {
			runLog["time_created"] = r.TimeCreated.String()
		}

		runLog["type"] = r.Type

		resources = append(resources, runLog)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataflowRunLogsDataSource().Schema["run_logs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("run_logs", resources); err != nil {
		return err
	}

	return nil
}
