// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetExportStatusDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetExportStatus,
		Schema: map[string]*schema.Schema{
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"latest_run_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_run": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_next_run": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsFleetExportStatus(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetExportStatusDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetExportStatusDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetExportStatusResponse
}

func (s *JmsFleetExportStatusDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetExportStatusDataSourceCrud) Get() error {
	request := oci_jms.GetExportStatusRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetExportStatus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetExportStatusDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetExportStatusDataSource-", JmsFleetExportStatusDataSource(), s.D))

	s.D.Set("latest_run_status", s.Res.LatestRunStatus)

	if s.Res.TimeLastRun != nil {
		s.D.Set("time_last_run", s.Res.TimeLastRun.String())
	}

	if s.Res.TimeNextRun != nil {
		s.D.Set("time_next_run", s.Res.TimeNextRun.String())
	}

	return nil
}
