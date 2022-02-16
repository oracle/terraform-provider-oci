// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v58/loganalytics"
)

func LogAnalyticsLogAnalyticsEntitiesSummaryDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsLogAnalyticsEntitiesSummary,
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
			"active_entities_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entities_with_has_logs_collected_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entities_with_management_agent_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsLogAnalyticsEntitiesSummary(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntitiesSummaryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsEntitiesSummaryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetLogAnalyticsEntitiesSummaryResponse
}

func (s *LogAnalyticsLogAnalyticsEntitiesSummaryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsEntitiesSummaryDataSourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsEntitiesSummaryRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetLogAnalyticsEntitiesSummary(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsEntitiesSummaryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsEntitiesSummaryDataSource-", LogAnalyticsLogAnalyticsEntitiesSummaryDataSource(), s.D))

	if s.Res.ActiveEntitiesCount != nil {
		s.D.Set("active_entities_count", *s.Res.ActiveEntitiesCount)
	}

	if s.Res.EntitiesWithHasLogsCollectedCount != nil {
		s.D.Set("entities_with_has_logs_collected_count", *s.Res.EntitiesWithHasLogsCollectedCount)
	}

	if s.Res.EntitiesWithManagementAgentCount != nil {
		s.D.Set("entities_with_management_agent_count", *s.Res.EntitiesWithManagementAgentCount)
	}

	return nil
}
