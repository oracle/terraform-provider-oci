// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationScheduledRunDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMeteringComputationScheduledRun,
		Schema: map[string]*schema.Schema{
			"scheduled_run_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularMeteringComputationScheduledRun(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationScheduledRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationScheduledRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.GetScheduledRunResponse
}

func (s *MeteringComputationScheduledRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationScheduledRunDataSourceCrud) Get() error {
	request := oci_metering_computation.GetScheduledRunRequest{}

	if scheduledRunId, ok := s.D.GetOkExists("scheduled_run_id"); ok {
		tmp := scheduledRunId.(string)
		request.ScheduledRunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.GetScheduledRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationScheduledRunDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ScheduleId != nil {
		s.D.Set("schedule_id", *s.Res.ScheduleId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	return nil
}
