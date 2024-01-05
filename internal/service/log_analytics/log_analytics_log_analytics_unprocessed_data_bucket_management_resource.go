// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement,
		Update:   updateLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement,
		Read:     readLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement,
		Delete:   deleteLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func updateLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func readLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteLogAnalyticsLogAnalyticsUnprocessedDataBucketManagement(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.UnprocessedDataBucket
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud) ID() string {
	return s.D.Get("namespace").(string)
}

func (s *LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud) Create() error {
	request := oci_log_analytics.SetUnprocessedDataBucketRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.SetUnprocessedDataBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UnprocessedDataBucket
	return nil
}

func (s *LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud) Update() error {
	request := oci_log_analytics.SetUnprocessedDataBucketRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.SetUnprocessedDataBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UnprocessedDataBucket
	return nil
}

func (s *LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud) Delete() error {
	request := oci_log_analytics.SetUnprocessedDataBucketRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	isEnabled := false
	request.IsEnabled = &isEnabled

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.SetUnprocessedDataBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UnprocessedDataBucket
	return nil
}

func (s *LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResourceCrud) SetData() error {
	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
