// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v49/loganalytics"
)

func init() {
	RegisterDatasource("oci_log_analytics_log_sets_count", LogAnalyticsLogSetsCountDataSource())
}

func LogAnalyticsLogSetsCountDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsLogSetsCount,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"log_sets_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsLogSetsCount(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogSetsCountDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).logAnalyticsClient()

	return ReadResource(sync)
}

type LogAnalyticsLogSetsCountDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetLogSetsCountResponse
}

func (s *LogAnalyticsLogSetsCountDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogSetsCountDataSourceCrud) Get() error {
	request := oci_log_analytics.GetLogSetsCountRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetLogSetsCount(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogSetsCountDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LogAnalyticsLogSetsCountDataSource-", LogAnalyticsLogSetsCountDataSource(), s.D))

	if s.Res.Count != nil {
		s.D.Set("log_sets_count", strconv.FormatInt(*s.Res.Count, 10))
	}

	return nil
}
