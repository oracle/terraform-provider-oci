// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"

	"github.com/oracle/terraform-provider-oci/crud"
)

func WorkRequestLogEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWorkRequestLogEntries,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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

func readWorkRequestLogEntries(d *schema.ResourceData, m interface{}) error {
	sync := &WorkRequestLogEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return crud.ReadResource(sync)
}

type WorkRequestLogEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListWorkRequestLogsResponse
}

func (s *WorkRequestLogEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WorkRequestLogEntriesDataSourceCrud) Get() error {
	request := oci_containerengine.ListWorkRequestLogsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.ListWorkRequestLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WorkRequestLogEntriesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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
		resources = ApplyFilters(f.(*schema.Set), resources, WorkRequestLogEntriesDataSource().Schema["work_request_log_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("work_request_log_entries", resources); err != nil {
		panic(err)
	}

	return
}
