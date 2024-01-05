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

func DataintegrationWorkspaceImportRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataintegrationWorkspaceImportRequests,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"import_status": {
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
			"import_request_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataintegrationWorkspaceImportRequestResource(),
						},
					},
				},
			},
		},
	}
}

func readDataintegrationWorkspaceImportRequests(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceImportRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

type DataintegrationWorkspaceImportRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataintegration.DataIntegrationClient
	Res    *oci_dataintegration.ListImportRequestsResponse
}

func (s *DataintegrationWorkspaceImportRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataintegrationWorkspaceImportRequestsDataSourceCrud) Get() error {
	request := oci_dataintegration.ListImportRequestsRequest{}

	if importStatus, ok := s.D.GetOkExists("import_status"); ok {
		request.ImportStatus = oci_dataintegration.ListImportRequestsImportStatusEnum(importStatus.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if projection, ok := s.D.GetOkExists("projection"); ok {
		request.Projection = oci_dataintegration.ListImportRequestsProjectionEnum(projection.(string))
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

	response, err := s.Client.ListImportRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListImportRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataintegrationWorkspaceImportRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataintegrationWorkspaceImportRequestsDataSource-", DataintegrationWorkspaceImportRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	workspaceImportRequest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ImportRequestSummaryToMap(item))
	}
	workspaceImportRequest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataintegrationWorkspaceImportRequestsDataSource().Schema["import_request_summary_collection"].Elem.(*schema.Resource).Schema)
		workspaceImportRequest["items"] = items
	}

	resources = append(resources, workspaceImportRequest)
	if err := s.D.Set("import_request_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
