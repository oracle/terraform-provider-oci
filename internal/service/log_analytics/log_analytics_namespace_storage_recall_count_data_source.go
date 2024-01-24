// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceStorageRecallCountDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsNamespaceStorageRecallCount,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"recall_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recall_failed": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recall_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recall_pending": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recall_succeeded": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsNamespaceStorageRecallCount(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageRecallCountDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceStorageRecallCountDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetRecallCountResponse
}

func (s *LogAnalyticsNamespaceStorageRecallCountDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceStorageRecallCountDataSourceCrud) Get() error {
	request := oci_log_analytics.GetRecallCountRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetRecallCount(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceStorageRecallCountDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceStorageRecallCountDataSource-", LogAnalyticsNamespaceStorageRecallCountDataSource(), s.D))

	if s.Res.RecallCount.RecallCount != nil {
		s.D.Set("recall_count", *s.Res.RecallCount.RecallCount)
	}

	if s.Res.RecallFailed != nil {
		s.D.Set("recall_failed", *s.Res.RecallFailed)
	}

	if s.Res.RecallLimit != nil {
		s.D.Set("recall_limit", *s.Res.RecallLimit)
	}

	if s.Res.RecallPending != nil {
		s.D.Set("recall_pending", *s.Res.RecallPending)
	}

	if s.Res.RecallSucceeded != nil {
		s.D.Set("recall_succeeded", *s.Res.RecallSucceeded)
	}

	return nil
}
