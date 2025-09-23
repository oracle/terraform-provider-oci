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

func JmsUtilsJavaMigrationAnalysiDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsUtilsJavaMigrationAnalysi,
		Schema: map[string]*schema.Schema{
			"java_migration_analysis_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"analysis_project_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"analysis_result_files": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"analysis_result_object_storage_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket": {
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
			"input_applications_object_storage_paths": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"metadata": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_jdk_version": {
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
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsUtilsJavaMigrationAnalysi(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsJavaMigrationAnalysiDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

type JmsUtilsJavaMigrationAnalysiDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_utils.JmsUtilsClient
	Res    *oci_jms_utils.GetJavaMigrationAnalysisResponse
}

func (s *JmsUtilsJavaMigrationAnalysiDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsUtilsJavaMigrationAnalysiDataSourceCrud) Get() error {
	request := oci_jms_utils.GetJavaMigrationAnalysisRequest{}

	if javaMigrationAnalysisId, ok := s.D.GetOkExists("java_migration_analysis_id"); ok {
		tmp := javaMigrationAnalysisId.(string)
		request.JavaMigrationAnalysisId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_utils")

	response, err := s.Client.GetJavaMigrationAnalysis(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsUtilsJavaMigrationAnalysiDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AnalysisProjectName != nil {
		s.D.Set("analysis_project_name", *s.Res.AnalysisProjectName)
	}

	s.D.Set("analysis_result_files", s.Res.AnalysisResultFiles)

	if s.Res.AnalysisResultObjectStoragePath != nil {
		s.D.Set("analysis_result_object_storage_path", *s.Res.AnalysisResultObjectStoragePath)
	}

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", []interface{}{PrincipalToMap(s.Res.CreatedBy)})
	} else {
		s.D.Set("created_by", nil)
	}

	s.D.Set("input_applications_object_storage_paths", s.Res.InputApplicationsObjectStoragePaths)

	if s.Res.Metadata != nil {
		s.D.Set("metadata", *s.Res.Metadata)
	}

	if s.Res.NamespaceName != nil {
		s.D.Set("namespace", *s.Res.NamespaceName)
	}

	if s.Res.TargetJdkVersion != nil {
		s.D.Set("target_jdk_version", *s.Res.TargetJdkVersion)
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

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}
