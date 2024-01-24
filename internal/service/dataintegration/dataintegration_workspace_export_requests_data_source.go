// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceExportRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataintegrationWorkspaceExportRequests,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"export_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"projection": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_ended_in_millis": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started_in_millis": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"export_request_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataintegrationWorkspaceExportRequestResource(),
						},
					},
				},
			},
		},
	}
}

func readDataintegrationWorkspaceExportRequests(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceExportRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceExportRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.ListExportRequestsResponse
}

func (s *DataintegrationWorkspaceExportRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceExportRequestsDataSourceCrud) Get() error {
	request := oci_dataintegration.ListExportRequestsRequest{}

	if exportStatus, ok := s.D.GetOkExists("export_status"); ok {
		request.ExportStatus = oci_dataintegration.ListExportRequestsExportStatusEnum(exportStatus.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if projection, ok := s.D.GetOkExists("projection"); ok {
		request.Projection = oci_dataintegration.ListExportRequestsProjectionEnum(projection.(string))
	}

	if timeEndedInMillis, ok := s.D.GetOkExists("time_ended_in_millis"); ok {
		tmp := timeEndedInMillis.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert timeEndedInMillis string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.TimeEndedInMillis = &tmpInt64
	}

	if timeStartedInMillis, ok := s.D.GetOkExists("time_started_in_millis"); ok {
		tmp := timeStartedInMillis.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert timeStartedInMillis string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.TimeStartedInMillis = &tmpInt64
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataintegration")

	response, err := s.Client.ListExportRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExportRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataintegrationWorkspaceExportRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceExportRequestsDataSource-", DataintegrationWorkspaceExportRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	workspaceExportRequest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ExportRequestSummaryToMap(item))
	}
	workspaceExportRequest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataintegrationWorkspaceExportRequestsDataSource().Schema["export_request_summary_collection"].Elem.(*schema.Resource).Schema)
		workspaceExportRequest["items"] = items
	}

	resources = append(resources, workspaceExportRequest)
	if err := s.D.Set("export_request_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
