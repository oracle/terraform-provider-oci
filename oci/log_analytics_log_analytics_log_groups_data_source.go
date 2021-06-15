// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v42/loganalytics"
)

func init() {
	RegisterDatasource("oci_log_analytics_log_analytics_log_groups", LogAnalyticsLogAnalyticsLogGroupsDataSource())
}

func LogAnalyticsLogAnalyticsLogGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsLogAnalyticsLogGroups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_analytics_log_group_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(LogAnalyticsLogAnalyticsLogGroupResource()),
						},
					},
				},
			},
		},
	}
}

func readLogAnalyticsLogAnalyticsLogGroups(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsLogGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).logAnalyticsClient()

	return ReadResource(sync)
}

type LogAnalyticsLogAnalyticsLogGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListLogAnalyticsLogGroupsResponse
}

func (s *LogAnalyticsLogAnalyticsLogGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsLogGroupsDataSourceCrud) Get() error {
	request := oci_log_analytics.ListLogAnalyticsLogGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListLogAnalyticsLogGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLogAnalyticsLogGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsLogGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LogAnalyticsLogAnalyticsLogGroupsDataSource-", LogAnalyticsLogAnalyticsLogGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	logAnalyticsLogGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsLogGroupSummaryToMap(item))
	}
	logAnalyticsLogGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsLogAnalyticsLogGroupsDataSource().Schema["log_analytics_log_group_summary_collection"].Elem.(*schema.Resource).Schema)
		logAnalyticsLogGroup["items"] = items
	}

	resources = append(resources, logAnalyticsLogGroup)
	if err := s.D.Set("log_analytics_log_group_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
