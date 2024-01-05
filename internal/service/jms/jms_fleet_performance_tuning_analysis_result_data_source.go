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

func JmsFleetPerformanceTuningAnalysisResultDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetPerformanceTuningAnalysisResult,
		Schema: map[string]*schema.Schema{
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"performance_tuning_analysis_result_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"application_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_installation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_installation_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"result": {
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
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"warning_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsFleetPerformanceTuningAnalysisResult(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetPerformanceTuningAnalysisResultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetPerformanceTuningAnalysisResultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetPerformanceTuningAnalysisResultResponse
}

func (s *JmsFleetPerformanceTuningAnalysisResultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetPerformanceTuningAnalysisResultDataSourceCrud) Get() error {
	request := oci_jms.GetPerformanceTuningAnalysisResultRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if performanceTuningAnalysisResultId, ok := s.D.GetOkExists("performance_tuning_analysis_result_id"); ok {
		tmp := performanceTuningAnalysisResultId.(string)
		request.PerformanceTuningAnalysisResultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetPerformanceTuningAnalysisResult(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetPerformanceTuningAnalysisResultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApplicationId != nil {
		s.D.Set("application_id", *s.Res.ApplicationId)
	}

	if s.Res.ApplicationInstallationId != nil {
		s.D.Set("application_installation_id", *s.Res.ApplicationInstallationId)
	}

	if s.Res.ApplicationInstallationPath != nil {
		s.D.Set("application_installation_path", *s.Res.ApplicationInstallationPath)
	}

	if s.Res.ApplicationName != nil {
		s.D.Set("application_name", *s.Res.ApplicationName)
	}

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.HostName != nil {
		s.D.Set("host_name", *s.Res.HostName)
	}

	if s.Res.ManagedInstanceId != nil {
		s.D.Set("managed_instance_id", *s.Res.ManagedInstanceId)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	s.D.Set("result", s.Res.Result)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.WarningCount != nil {
		s.D.Set("warning_count", *s.Res.WarningCount)
	}

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}
