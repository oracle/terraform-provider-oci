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

func LogAnalyticsNamespaceRulesSummaryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsNamespaceRulesSummary,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"ingest_time_rules_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"saved_search_rules_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsNamespaceRulesSummary(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceRulesSummaryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceRulesSummaryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetRulesSummaryResponse
}

func (s *LogAnalyticsNamespaceRulesSummaryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceRulesSummaryDataSourceCrud) Get() error {
	request := oci_log_analytics.GetRulesSummaryRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetRulesSummary(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceRulesSummaryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceRulesSummaryDataSource-", LogAnalyticsNamespaceRulesSummaryDataSource(), s.D))

	if s.Res.IngestTimeRulesCount != nil {
		s.D.Set("ingest_time_rules_count", *s.Res.IngestTimeRulesCount)
	}

	if s.Res.SavedSearchRulesCount != nil {
		s.D.Set("saved_search_rules_count", *s.Res.SavedSearchRulesCount)
	}

	if s.Res.TotalCount != nil {
		s.D.Set("total_count", *s.Res.TotalCount)
	}

	return nil
}
