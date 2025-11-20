// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psa

import (
	"context"

	"github.com/oracle/oci-go-sdk/v65/psa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func PsaWorkRequestLogEntriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readPsaWorkRequestLogEntries,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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

func readPsaWorkRequestLogEntries(d *schema.ResourceData, m interface{}) error {
	sync := &PsaWorkRequestLogEntriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivateServiceAccessClient()

	return tfresource.ReadResource(sync)
}

type PsaWorkRequestLogEntriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *psa.PrivateServiceAccessClient
	Res    *psa.ListPsaWorkRequestLogsResponse
}

func (s *PsaWorkRequestLogEntriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsaWorkRequestLogEntriesDataSourceCrud) Get() error {
	request := psa.ListPsaWorkRequestLogsRequest{}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psa")

	response, err := s.Client.ListPsaWorkRequestLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsaWorkRequestLogEntriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("PsaWorkRequestLogEntriesDataSource-", PsaWorkRequestLogEntriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		workRequestLogEntry := map[string]interface{}{}

		if r.Message != nil {
			workRequestLogEntry["message"] = *r.Message
		}

		if r.Timestamp != nil {
			workRequestLogEntry["timestamp"] = r.Timestamp.String()
		}

		resources = append(resources, workRequestLogEntry)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, PsaWorkRequestLogEntriesDataSource().Schema["work_request_log_entries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("work_request_log_entries", resources); err != nil {
		return err
	}

	return nil
}
