// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v56/containerengine"
)

func ContainerengineWorkRequestLogEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineWorkRequestLogEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"work_request_log_entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readContainerengineWorkRequestLogEntries(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineWorkRequestLogEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineWorkRequestLogEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListWorkRequestLogsResponse
}

func (s *ContainerengineWorkRequestLogEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineWorkRequestLogEntriesDataSourceCrud) Get() error {
	request := oci_containerengine.ListWorkRequestLogsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListWorkRequestLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineWorkRequestLogEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineWorkRequestLogEntriesDataSource-", ContainerengineWorkRequestLogEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		workRequestLogEntry := map[string]interface{}{}

		if r.Message != nil {
			workRequestLogEntry["message"] = *r.Message
		}

		if r.Timestamp != nil {
			workRequestLogEntry["timestamp"] = *r.Timestamp
		}

		resources = append(resources, workRequestLogEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineWorkRequestLogEntriesDataSource().Schema["work_request_log_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("work_request_log_entries", resources); err != nil {
		return err
	}

	return nil
}
