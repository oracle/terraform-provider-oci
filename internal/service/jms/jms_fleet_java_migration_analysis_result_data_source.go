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

func JmsFleetJavaMigrationAnalysisResultDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsFleetJavaMigrationAnalysisResult,
		Schema: map[string]*schema.Schema{
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"java_migration_analysis_result_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"application_execution_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_path": {
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
			"metadata": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"object_storage_upload_dir_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_jdk_version": {
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
			"work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularJmsFleetJavaMigrationAnalysisResult(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetJavaMigrationAnalysisResultDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetJavaMigrationAnalysisResultDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.GetJavaMigrationAnalysisResultResponse
}

func (s *JmsFleetJavaMigrationAnalysisResultDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetJavaMigrationAnalysisResultDataSourceCrud) Get() error {
	request := oci_jms.GetJavaMigrationAnalysisResultRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if javaMigrationAnalysisResultId, ok := s.D.GetOkExists("java_migration_analysis_result_id"); ok {
		tmp := javaMigrationAnalysisResultId.(string)
		request.JavaMigrationAnalysisResultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.GetJavaMigrationAnalysisResult(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsFleetJavaMigrationAnalysisResultDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("application_execution_type", s.Res.ApplicationExecutionType)

	if s.Res.ApplicationKey != nil {
		s.D.Set("application_key", *s.Res.ApplicationKey)
	}

	if s.Res.ApplicationName != nil {
		s.D.Set("application_name", *s.Res.ApplicationName)
	}

	if s.Res.ApplicationPath != nil {
		s.D.Set("application_path", *s.Res.ApplicationPath)
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

	if s.Res.Metadata != nil {
		s.D.Set("metadata", *s.Res.Metadata)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	s.D.Set("object_list", s.Res.ObjectList)

	if s.Res.ObjectStorageUploadDirPath != nil {
		s.D.Set("object_storage_upload_dir_path", *s.Res.ObjectStorageUploadDirPath)
	}

	if s.Res.SourceJdkVersion != nil {
		s.D.Set("source_jdk_version", *s.Res.SourceJdkVersion)
	}

	if s.Res.TargetJdkVersion != nil {
		s.D.Set("target_jdk_version", *s.Res.TargetJdkVersion)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}
