// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_utils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_utils "github.com/oracle/oci-go-sdk/v65/jmsutils"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsUtilsPerformanceTuningAnalysiDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsUtilsPerformanceTuningAnalysi,
		Schema: map[string]*schema.Schema{
			"performance_tuning_analysis_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"analysis_project_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"artifact_object_storage_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"result": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"result_object_storage_path": {
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

func readSingularJmsUtilsPerformanceTuningAnalysi(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsPerformanceTuningAnalysiDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

type JmsUtilsPerformanceTuningAnalysiDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_utils.JmsUtilsClient
	Res    *oci_jms_utils.GetPerformanceTuningAnalysisResponse
}

func (s *JmsUtilsPerformanceTuningAnalysiDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsUtilsPerformanceTuningAnalysiDataSourceCrud) Get() error {
	request := oci_jms_utils.GetPerformanceTuningAnalysisRequest{}

	if performanceTuningAnalysisId, ok := s.D.GetOkExists("performance_tuning_analysis_id"); ok {
		tmp := performanceTuningAnalysisId.(string)
		request.PerformanceTuningAnalysisId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_utils")

	response, err := s.Client.GetPerformanceTuningAnalysis(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsUtilsPerformanceTuningAnalysiDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AnalysisProjectName != nil {
		s.D.Set("analysis_project_name", *s.Res.AnalysisProjectName)
	}

	if s.Res.ArtifactObjectStoragePath != nil {
		s.D.Set("artifact_object_storage_path", *s.Res.ArtifactObjectStoragePath)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", []interface{}{PrincipalToMap(s.Res.CreatedBy)})
	} else {
		s.D.Set("created_by", nil)
	}

	s.D.Set("result", s.Res.Result)

	if s.Res.ResultObjectStoragePath != nil {
		s.D.Set("result_object_storage_path", *s.Res.ResultObjectStoragePath)
	}

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
